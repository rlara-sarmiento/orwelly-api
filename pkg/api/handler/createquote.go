package handler

import (
	"io"
	"net/http"
)

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "I'll create a quote here\n")
}