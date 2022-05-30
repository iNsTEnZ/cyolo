package rest

import (
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (api *API) process(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	api.rateLimiter <- struct{}{}
	defer func() { <-api.rateLimiter }()

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println("error reading request body")
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	JSON(writer, api.service.Process(string(body)))
}

func (api *API) histogram(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	api.rateLimiter <- struct{}{}
	defer func() { <-api.rateLimiter }()

	value := req.URL.Query().Get("top")

	if value == "" {
		value = "5"
	}

	if top, err := strconv.Atoi(value); err == nil {
		result := api.service.Histogram(top)
		TEXT(writer, api.printer.Print(result))
	} else {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
