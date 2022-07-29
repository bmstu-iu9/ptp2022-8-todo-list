package router

import "github.com/julienschmidt/httprouter"

// New returns preconfigured httprouter.Router.
func New() *httprouter.Router {
	return &httprouter.Router{
		HandleMethodNotAllowed: true,
	}
}
