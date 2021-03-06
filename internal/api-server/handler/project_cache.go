package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
	"time"

	"github.com/Donders-Institute/filer-gateway/internal/api-server/config"
	"github.com/Donders-Institute/filer-gateway/internal/task"
	"github.com/Donders-Institute/filer-gateway/pkg/swagger/server/models"
	fp "github.com/Donders-Institute/tg-toolset-golang/pkg/filepath"
	log "github.com/Donders-Institute/tg-toolset-golang/pkg/logger"
	"github.com/Donders-Institute/tg-toolset-golang/project/pkg/acl"
	"github.com/go-redis/redis/v8"
)

type projectResource struct {
	storage *models.StorageResponse
	members []*models.Member
}

// ProjectResourceCache is an in-memory store for caching `projectResource` of all existing projects
// on the filer.
type ProjectResourceCache struct {

	// Config is the general API server configuration.
	Config config.Configuration

	// Context is the API server context.
	Context context.Context

	// Notifier is the redis channel subscription via which
	// a refresh on a given project can be triggered on-demand.
	Notifier <-chan *redis.Message

	// IsStopped indicates whether the cache service is stopped.
	IsStopped bool
	store     map[string]*projectResource
	mutex     sync.Mutex
}

// init initializes the cache with first reload.
func (c *ProjectResourceCache) Init() {

	c.IsStopped = false

	// first refresh
	c.refresh()

	// every 10 minutes??
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Infof("refreshing project cache")
				c.refresh()
				log.Infof("project cache refreshed")
			case m := <-c.Notifier:
				// interpret request payload
				p := task.UpdateProjectPayload{}
				if err := json.Unmarshal([]byte(m.Payload), &p); err != nil {
					log.Errorf("unknown update payload: %s", m.Payload)
					continue
				}
				// perform cache update
				if _, err := c.getResource(p.ProjectID, true); err != nil {
					log.Errorf("fail to update cache for project %s: %s", p.ProjectID, err)
					continue
				}
				log.Infof("cache updated for project: %s", p.ProjectID)

			case <-c.Context.Done():
				log.Infof("project cache refresh stopped")
				c.IsStopped = true
				return
			}
		}
	}()

	log.Infof("project cache initalized")
}

// refresh update the cache with up-to-data project resources.
func (c *ProjectResourceCache) refresh() {

	nworkers := runtime.NumCPU()

	pnumbers := make(chan string, nworkers*2)
	resources := make(chan struct {
		pnumber  string
		resource *projectResource
	})

	wg := sync.WaitGroup{}
	// start concurrent workers to get project resources from the filer.
	for i := 0; i < nworkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for pnumber := range pnumbers {
				storage, members, err := getProjectResource(pnumber, c.Config)
				if err != nil {
					log.Errorf("cannot get filer resource for %s: %s", pnumber, err)
				}
				resources <- struct {
					pnumber  string
					resource *projectResource
				}{
					pnumber,
					&projectResource{
						storage: storage,
						members: members,
					},
				}
			}
		}()
	}

	// go routine to list all directories in the /project folder
	go func(path string) {
		// close the dirs channel on exit
		defer close(pnumbers)

		objs, err := fp.ListDir(path)
		if err != nil {
			log.Errorf("cannot get content of path: %s", path)
			return
		}

		for _, obj := range objs {
			pnumbers <- filepath.Base(obj)
		}

	}(PathProject)

	// go routine to wait for all workers to complete and close the resources channel.
	go func() {
		wg.Wait()
		close(resources)
	}()

	// new data map
	d := make(map[string]*projectResource)

	// merge resources into internal store
	for r := range resources {
		d[r.pnumber] = r.resource
	}

	// set store to new data map
	c.mutex.Lock()
	c.store = d
	c.mutex.Unlock()
}

// getResource finds and returns project resource from the cache.
// An error is returned if the project doesn't exist in cache.
func (c *ProjectResourceCache) getResource(pnumber string, force bool) (*projectResource, error) {
	if r, ok := c.store[pnumber]; !ok || force {

		// not found in cache, try fetch from the filer.
		storage, members, err := getProjectResource(pnumber, c.Config)
		if err != nil {
			return nil, err
		}

		// update cache with data retrieved from the filer.
		c.mutex.Lock()
		c.store[pnumber] = &projectResource{
			storage: storage,
			members: members,
		}
		c.mutex.Unlock()

		// return the data from cache
		return c.store[pnumber], nil
	} else {
		return r, nil
	}
}

// getProjectResource retrieves storage resource and access roles of a given project.
func getProjectResource(pnumber string, cfg config.Configuration) (*models.StorageResponse, []*models.Member, error) {

	path, err := pid2path(pnumber)

	if err != nil {
		return nil, nil, &ResponseError{code: 404, err: err.Error()}
	}

	// Get Storage Resource
	system, quota, usage, err := getStorageQuota(cfg, path)
	// Return response error based on error code.
	if err != nil {
		return nil, nil, err
	}

	members, err := getMemberRoles(path)
	if err != nil {
		return nil, nil, err
	}

	quotaGb := quota >> 30
	usageMb := usage >> 20
	storage := &models.StorageResponse{
		QuotaGb: &quotaGb,
		System:  &system,
		UsageMb: &usageMb,
	}

	return storage, members, nil
}

// getMemberRoles retrives member roles applied on the path.
func getMemberRoles(path string) ([]*models.Member, error) {

	members := make([]*models.Member, 0)

	runner := acl.Runner{
		RootPath:   path,
		FollowLink: true,
		SkipFiles:  true,
		Nthreads:   1,
	}

	chanOut, err := runner.GetRoles(false)

	// we know it's path not found error because this is the only case the runner.GetRoles returns an error.
	// TODO: maybe the runner should return an explicit error type.
	if err != nil {
		return members, &ResponseError{code: 500, err: fmt.Sprintf("cannot get role: %s", path)}
	}

	// only one object is expected from the channel as the recursion is disabled on the runner function.
	for o := range chanOut {
		log.Debugf("found project memebers on %s, %+v\n", o.Path, o.RoleMap)
		for r, users := range o.RoleMap {
			// exclude the system role.
			if r == acl.System {
				continue
			}
			rname := r.String()
			for i := range users {
				m := models.Member{
					UserID: &users[i],
					Role:   &rname,
				}
				members = append(members, &m)
			}
		}
	}

	return members, nil
}

// pid2path converts project id to file system path.
func pid2path(pid string) (string, error) {
	var path string
	if matched, _ := regexp.MatchString("^[0-9]{7,}", pid); matched {
		// input pid is a project number
		path = filepath.Join(PathProject, pid)
	} else {
		return path, fmt.Errorf("invalid project id: %s", pid)
	}

	// evaluate symlink to its absolute path.
	return filepath.EvalSymlinks(path)
}
