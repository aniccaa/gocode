package httputils

import "net/http"

type APIFunc func(w http.ResponseWriter, r *http.Request)
