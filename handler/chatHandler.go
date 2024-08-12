package handler

import (
	"net/http"
	"strconv"

	db "github.com/keiko30/chatbot/db"
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
	id, err := strconv.ParseUint(message, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	chat := db.GetQuestion(uint(id))
	view.Message(chat).Render(r.Context(), w)
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	chats, err := db.GetQuestions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	view.Questions(chats).Render(r.Context(), w)
}
