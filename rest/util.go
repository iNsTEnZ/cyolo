package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)

	if err != nil {
		log.Printf("resp: %v", data)
	}
}

func TEXT(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(data))

	if err != nil {
		log.Printf("resp: %v", data)
	}
}
