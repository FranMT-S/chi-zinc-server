package routes

import (
	"github.com/go-chi/chi/v5"
)

func ApiV1Router() *chi.Mux {
	router := chi.NewMux()

	router.Route("/api", func(r chi.Router) {
		r.Mount("/mails", MailRouter())
	})

	return router
}
