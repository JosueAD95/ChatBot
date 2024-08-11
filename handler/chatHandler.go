package handler

import (
	"fmt"
	"net/http"

	view "github.com/keiko30/chatbot/view"
)

func Index(w http.ResponseWriter, r *http.Request) {
	view.Index().Render(r.Context(), w)
}

func ResponseMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.PostFormValue("message"))
}
