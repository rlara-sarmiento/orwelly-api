package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rlara-sarmiento/orwelly-api/pkg/api/core"
	"github.com/rlara-sarmiento/orwelly-api/pkg/api/model"
)

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	var req createQuoteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendErrResponse(w, http.StatusBadRequest, "Cannot parse request")
		return
	}

	quote, err := core.Create(req.toQuote())
	if err != nil {
		sendErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusCreated, createQuoteResponseFromQuote(quote))
}

type createQuoteRequest struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

func (r createQuoteRequest) toQuote() model.Quote {
	return model.NewQuote(r.Author, r.Text)
}

type createQuoteResponse struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

func createQuoteResponseFromQuote(quote model.Quote) createQuoteResponse {
	return createQuoteResponse{
		Id:     quote.Id,
		Author: quote.Author,
		Text:   quote.Text,
	}
}
