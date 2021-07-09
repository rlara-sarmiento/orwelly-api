package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rlara-sarmiento/orwelly-api/pkg/api/core"
)

func ListQuotes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello! This is orwelly-api list\n")
	y := core.ConsultList()
	str := fmt.Sprintf("%v", y)
	io.WriteString(w, str)
}
func aux(w http.ResponseWriter, r *http.Request) {
	//x := core.ConsultList()
	//for i := range x {
	//	sendJsonResponse(w, i, x[i])
	//}
}
