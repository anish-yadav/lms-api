package webresponse

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Response struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

func RespondWithError(w http.ResponseWriter, status int, message interface{}) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	response := &Response{
		Error: message,
		Data:  nil,
	}
	bin, _ := json.Marshal(response)
	if _, err := w.Write(bin); err != nil {
		log.Errorf("webservice.RespondWriter: failed to write data : %s", err.Error())
	}
}


func RespondWithSuccess(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-type", "application/json")
	response := &Response{
		Error: nil,
		Data: data,
	}
	bin, _ := json.Marshal(response)
	if _, err := w.Write(bin); err != nil {
		log.Errorf("webservice.RespondWriter: failed to write data : %s", err.Error())
	}
}
