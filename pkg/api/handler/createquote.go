package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rlara-sarmiento/orwelly-api/pkg/api/core"
	"github.com/rlara-sarmiento/orwelly-api/pkg/api/model"
)

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	var req CreateQuoteRequest
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

	sendJsonResponse(w, http.StatusOK, createQuoteResponseFromQuote(quote))
}

type CreateQuoteRequest struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

func (r CreateQuoteRequest) toQuote() model.Quote {
	return model.NewQuote(r.Author, r.Text)
}

type CreateQuoteResponse struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

func createQuoteResponseFromQuote(quote model.Quote) CreateQuoteResponse {
	return CreateQuoteResponse{
		Id:     quote.Id,
		Author: quote.Author,
		Text:   quote.Text,
	}
}
