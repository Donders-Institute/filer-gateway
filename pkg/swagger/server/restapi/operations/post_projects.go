// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostProjectsHandlerFunc turns a function with the right signature into a post projects handler
type PostProjectsHandlerFunc func(PostProjectsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostProjectsHandlerFunc) Handle(params PostProjectsParams) middleware.Responder {
	return fn(params)
}

// PostProjectsHandler interface for that can handle valid post projects params
type PostProjectsHandler interface {
	Handle(PostProjectsParams) middleware.Responder
}

// NewPostProjects creates a new http.Handler for the post projects operation
func NewPostProjects(ctx *middleware.Context, handler PostProjectsHandler) *PostProjects {
	return &PostProjects{Context: ctx, Handler: handler}
}

/*PostProjects swagger:route POST /projects postProjects

provision filer resource for a new project.

*/
type PostProjects struct {
	Context *middleware.Context
	Handler PostProjectsHandler
}

func (o *PostProjects) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostProjectsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
