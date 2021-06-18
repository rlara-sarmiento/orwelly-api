package handler

import (
	"encoding/json"
	"net/http"
)

func sendJsonResponse(w http.ResponseWriter, code int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func sendErrResponse(w http.ResponseWriter, code int, msg string) {
	err := newErr(msg)
	sendJsonResponse(w, code, err)
}

type Err struct {
	Reason string `json:"reason"`
}

func newErr(msg string) Err {
	return Err{
		Reason: msg,
	}
}
