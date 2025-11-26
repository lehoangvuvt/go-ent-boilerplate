package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func StartServer(r *chi.Mux) {
	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
