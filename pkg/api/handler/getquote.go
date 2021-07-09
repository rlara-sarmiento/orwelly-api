package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rlara-sarmiento/orwelly-api/pkg/api/core"
	"github.com/rlara-sarmiento/orwelly-api/pkg/api/model"
)

func GetQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	quote, err := core.Get(id)

	if err == errors.New("not found") {
		NotFound(w, r)
		return
	}

	sendJsonResponse(w, http.StatusOK, getQuoteResponseFromQuote(quote))
}

type getQuoteResponse struct {
	Id     int
	Author string
	Text   string
}

func getQuoteResponseFromQuote(quote model.Quote) getQuoteResponse {
	return getQuoteResponse{
		Id:     quote.Id,
		Author: quote.Author,
		Text:   quote.Text,
	}
}
