package items

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers(mux *httprouter.Router, service Service, logger *log.Logger) {
	res := resource{service, logger}

	mux.GET("/users/:userId/items", res.handleGetAll)
	mux.GET("/users/:userId/items/:itemId", res.handleGetOne)
	mux.PUT("/users/:userId/items/:itemId", res.handlePut)
	mux.PATCH("/users/:userId/items/:itemId", res.handlePatch)
}

type resource struct {
	service Service
	logger  *log.Logger
}

func (res *resource) handleGetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	items, err := res.service.GetAll()
	if err != nil {
		res.logger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		res.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (res *resource) handleGetOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (res *resource) handlePut(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func (res *resource) handlePatch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func getUserId(p httprouter.Params) (int, error) {
	return strconv.Atoi(p.ByName("userId"))
}

func getItemId(p httprouter.Params) (int, error) {
	return strconv.Atoi(p.ByName("userId"))
}
