package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rlara-sarmiento/orwelly-api/pkg/api/core"
	"github.com/rlara-sarmiento/orwelly-api/pkg/api/datastore"
	"github.com/rlara-sarmiento/orwelly-api/pkg/api/model"
)

var ds datastore.Quotes

func GetQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req model.Quote
	id, _ := strconv.Atoi(vars["id"])
	x := core.ConsultList()
	y, err := x.Get(id)
	fmt.Println(y, err)
	if err == errors.New("Unimplemented") {
		NotFound(w, r)
		return
	}
	req.Id, req.Author, req.Text = id, y.Author, y.Text
	sendJsonResponse(w, http.StatusOK, req)
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
