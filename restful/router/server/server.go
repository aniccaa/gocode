package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"jhana/restful/router"
	"net/http"
)

type serverRouter struct {
	backend Backend
	routes  []router.Route
}

func (r *serverRouter) Routes() []router.Route {
	return r.routes
}

func NewRouter() router.Router {
	s := &serverRouter{}
	s.initRouter()
	return s
}

func (r *serverRouter) initRouter() []router.Route {
	r.routes = []router.Route{
		router.NewGetRouter("/server/info", r.getServerInfo),
		router.NewPostRouter("/server/info", r.postServerInfo),
	}
	return r.routes
}

func (r *serverRouter) getServerInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "server info")
}

func (r *serverRouter) postServerInfo(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("[server] post server info failed. read all body error")
		return
	}
	fmt.Fprint(w, string(reqBody))
}
