package service

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ServiceData struct {
	Result interface{} `json:"result"`
}

type RequestError struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, httpCode int, err error) {
	WriteJSON(w, httpCode, RequestError{Error: err.Error()})
}

func WriteData(w http.ResponseWriter, httpCode int, d interface{}) {
	WriteJSON(w, httpCode, ServiceData{Result: strconv.FormatUint(d.(uint64), 10)})
}

func WriteJSON(w http.ResponseWriter, httpCode int, d interface{}) {
	j, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, _ = w.Write(j)
}
