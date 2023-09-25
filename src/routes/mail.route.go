package routes

import (
	"github.com/FranMT-S/chi-zinc-server/src/controller"
	"github.com/go-chi/chi/v5"
)

// a router returns with the endpoints to make email requests
func MailRouter() *chi.Mux {
	router := chi.NewMux()

	router.Get("/{id}", controller.GetMail)
	router.Get("/from-{from}-max-{max}", controller.GetAllMailsSummary)
	router.Get("/from-{from}-max-{max}-terms-{terms}", controller.FindMailsSummary)

	return router
}
