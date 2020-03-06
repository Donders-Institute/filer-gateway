package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Donders-Institute/filer-gateway/internal/handlers"
	"github.com/Donders-Institute/filer-gateway/pkg/swagger/server/restapi"
	"github.com/Donders-Institute/filer-gateway/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/loads"

	log "github.com/sirupsen/logrus"
)

var (
	//optsConfig  *string
	optsVerbose *bool
)

func init() {
	//optsConfig = flag.String("c", "config.yml", "set the `path` of the configuration file")
	optsVerbose = flag.Bool("v", false, "print debug messages")

	flag.Usage = usage

	flag.Parse()

	// set logging
	log.SetOutput(os.Stderr)

	// set logging level
	llevel := log.InfoLevel
	if *optsVerbose {
		llevel = log.DebugLevel
	}
	log.SetLevel(llevel)
}

func usage() {
	fmt.Printf("\nAPI server for filer gateway\n")
	fmt.Printf("\nUSAGE: %s [OPTIONS]\n", os.Args[0])
	fmt.Printf("\nOPTIONS:\n")
	flag.PrintDefaults()
	fmt.Printf("\n")
}

func main() {
	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewFilerGatewayAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	// associate handlers with implementations
	api.PostProjectsHandler = operations.PostProjectsHandlerFunc(handlers.CreateProject())
	api.PatchProjectsIDHandler = operations.PatchProjectsIDHandlerFunc(handlers.UpdateProject())
	api.GetProjectsIDHandler = operations.GetProjectsIDHandlerFunc(handlers.GetProjectResource())
	api.GetProjectsIDMembersHandler = operations.GetProjectsIDMembersHandlerFunc(handlers.GetProjectMembers())
	api.GetProjectsIDStorageHandler = operations.GetProjectsIDStorageHandlerFunc(handlers.GetProjectStorage())

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
