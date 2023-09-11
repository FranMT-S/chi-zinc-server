package routes

import (
	"github.com/FranMT-S/chi-zinc-server/src/controller"
	"github.com/go-chi/chi/v5"
)

func MailRouter() *chi.Mux {
	router := chi.NewMux()

	router.Get("/", controller.GetTotalMail)
	router.Get("/all-sumary/from-{from}-max-{max}", controller.GetAllMailsSummary)
	router.Get("/mail/{id}", controller.GetMail)
	router.Post("/find-mails", controller.FindMailsSummary)

	return router
}
