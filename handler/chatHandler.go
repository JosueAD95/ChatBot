package handler

import (
	"net/http"

	view "github.com/keiko30/chatbot/view"
)

func Index(w http.ResponseWriter, r *http.Request) {
	view.Index().Render(r.Context(), w)
}

func ResponseMessage(w http.ResponseWriter, r *http.Request) {
	message := r.PostFormValue("message")
	if message == "" {
		w.Header().Set("x-missing-field", "message")
		w.WriteHeader(http.StatusLengthRequired)
		return
	}
	view.Message(message, "Test answer").Render(r.Context(), w)
}
