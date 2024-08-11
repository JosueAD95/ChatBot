package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	handler "github.com/keiko30/chatbot/handler"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", handler.Index)
	router.Post("/sendChatMessage/", handler.ResponseMessage)

	listeAddr := os.Getenv("LISTEN_ADDRESS")
	slog.Info("HTTP server started", "Listening address", listeAddr)
	http.ListenAndServe(listeAddr, router)
}
