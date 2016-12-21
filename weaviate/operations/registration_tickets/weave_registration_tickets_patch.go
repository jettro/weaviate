package registration_tickets


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveRegistrationTicketsPatchHandlerFunc turns a function with the right signature into a weave registration tickets patch handler
type WeaveRegistrationTicketsPatchHandlerFunc func(WeaveRegistrationTicketsPatchParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveRegistrationTicketsPatchHandlerFunc) Handle(params WeaveRegistrationTicketsPatchParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveRegistrationTicketsPatchHandler interface for that can handle valid weave registration tickets patch params
type WeaveRegistrationTicketsPatchHandler interface {
	Handle(WeaveRegistrationTicketsPatchParams, interface{}) middleware.Responder
}

// NewWeaveRegistrationTicketsPatch creates a new http.Handler for the weave registration tickets patch operation
func NewWeaveRegistrationTicketsPatch(ctx *middleware.Context, handler WeaveRegistrationTicketsPatchHandler) *WeaveRegistrationTicketsPatch {
	return &WeaveRegistrationTicketsPatch{Context: ctx, Handler: handler}
}

/*WeaveRegistrationTicketsPatch swagger:route PATCH /registrationTickets/{registrationTicketId} registrationTickets weaveRegistrationTicketsPatch

Updates an existing registration ticket. This method supports patch semantics.

*/
type WeaveRegistrationTicketsPatch struct {
	Context *middleware.Context
	Handler WeaveRegistrationTicketsPatchHandler
}

func (o *WeaveRegistrationTicketsPatch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveRegistrationTicketsPatchParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}