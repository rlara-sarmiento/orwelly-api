package handler

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello! This is orwelly-api\n")
}
