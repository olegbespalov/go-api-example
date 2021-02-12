package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Message writes a message to the http response
func Message(w http.ResponseWriter, code int, message string) {
	resp := &response{
		Code:    code,
		Message: message,
	}

	Dump(w, code, resp)
}

// Dump writes an interface as json
func Dump(w http.ResponseWriter, code int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	b, err := json.Marshal(resp)
	if err != nil {
		log.Println(err.Error())
	}

	w.Write(b)
}
