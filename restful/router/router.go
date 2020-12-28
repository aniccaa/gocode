package router

import "jhana/restful/httputils"

type Router interface {
	Routes() []Route
}

type Route interface {
	Method() string
	Path() string
	Handler() httputils.APIFunc
}
