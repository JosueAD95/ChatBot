package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	data "github.com/keiko30/chatbot/db"
	handler "github.com/keiko30/chatbot/handler"
)

func main() {
	listeAddr := setup()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", handler.Index)
	router.Get("/getQuestions/", handler.GetQuestions)
	router.Post("/sendChatMessage/", handler.ResponseMessage)

	slog.Info("HTTP server started", "Listening address", listeAddr)
	http.ListenAndServe(listeAddr, router)
}

func setup() string {
	if err := godotenv.Load(); err != nil {
		panic("Local variable couldn't be loaded.")
	}

	data.ConnectDB()
	if os.Getenv("MIGRATE_DB") == "Yes" {
		data.MigrateChatDB()
	}

	if os.Getenv("FILL_DB") == "Yes" {
		data.CreateChats()
	}

	return os.Getenv("LISTEN_ADDRESS")
}
