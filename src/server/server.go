package myServer

import (
	"fmt"
	"net/http"
	"os"

	myMiddleware "github.com/FranMT-S/chi-zinc-server/src/middleware"
	"github.com/FranMT-S/chi-zinc-server/src/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Iserve interface {
	Start()
}

var _server Iserve

type server struct {
	router *chi.Mux
}

// Server returns a single server instance
func Server() Iserve {

	if _server == nil {
		_server = &server{router: chi.NewRouter()}
	}

	return _server
}

// Start initialize the server
func (_server *server) Start() {

	_server.mountHandlers()

	fmt.Printf("Server listening on port %s...\n", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), _server.router)

}

func (s *server) mountHandlers() {
	// Mount all Middleware here
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.CleanPath)
	s.router.Use(middleware.AllowContentType("application/json"))
	s.router.Use(cors.Handler(cors.Options{

		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	s.router.Use(myMiddleware.JsonMiddleware)
	s.router.Use(myMiddleware.LogBookMiddleware)

	s.router.Route("/", func(r chi.Router) {
		r.Get("/", func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte("Hello World!"))
		})

		r.Mount("/v1", routes.ApiV1Router())
	})

}
