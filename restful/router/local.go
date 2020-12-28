package router

import (
	"jhana/restful/httputils"
)

type localRoute struct {
	method  string
	path    string
	handler httputils.APIFunc
}

func (l localRoute) Method() string {
	return l.method
}

func (l localRoute) Path() string {
	return l.path
}

func (l localRoute) Handler() httputils.APIFunc {
	return l.handler
}

func NewRouter(method string, path string, handler httputils.APIFunc) Route {
	r := localRoute{method: method, path: path, handler: handler}
	return r
}

func NewGetRouter(path string, handler httputils.APIFunc) Route {
	return NewRouter("Get", path, handler)
}

func NewPostRouter(path string, handler httputils.APIFunc) Route {
	return NewRouter("Post", path, handler)
}

func NewPatchRouter(path string, handler httputils.APIFunc) Route {
	return NewRouter("Patch", path, handler)
}
