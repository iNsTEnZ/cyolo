package rest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (api *API) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	JSON(w, map[string]interface{}{
		"status": "ok",
	})
}
