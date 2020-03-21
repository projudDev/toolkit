package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func SendError(w http.ResponseWriter, statusCode int, data interface{}) {
	SendJSON(w, statusCode, data)
}
