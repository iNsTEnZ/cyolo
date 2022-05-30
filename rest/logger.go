package rest

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (api *API) RequestLogger(next httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		next(w, r, ps)
		log.Printf("method: %s, uri: %s, name: %s", r.Method, r.RequestURI, name)
	}
}
