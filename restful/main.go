package main

import (
	"fmt"
	"jhana/restful/router/server"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:1228")
	if err != nil {
		log.Error("listen error")
		return
	}

	//
	server.NewRouter()

	if err = http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "welcome to homepage")
	})); err != nil {
		log.Error("serve error")
	}
}
