package items

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers(mux *httprouter.Router, logger *log.Logger) {
	res := resource{logger}

	mux.GET("/users/:id/items", res.handleGetAll)
	mux.GET("/users/:userId/items/:itemId", res.handleGetOne)
	mux.PUT("/users/:userId/items/:itemId", res.handlePut)
	mux.PATCH("/users/:userId/items/:itemId", res.handlePatch)
}

type resource struct {
	//service Service
	logger *log.Logger
}

func (res *resource) handleGetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (res *resource) handleGetOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (res *resource) handlePut(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

