package handlers

import (
	"go-rest-api/internal/middleware"
	"net/http"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

/*
 ! Middleware basically are handlers between request and response

request --> Middlewares... --> Handler


*/

// Starts with Capital means it can be used in other
// Capital suggests Public
func Handler(r *chi.Mux) {
	// Global Middleware
	logrus.Info("Middleware: Stripping slashes")
	r.Use(chimiddle.StripSlashes) // strips slashed from the query endpoint


	// This is a route
	r.Route("/employee", func(router chi.Router) {
		// Middleware
		logrus.Debug("Middleware: Running Authorization Middleware")
		// router.Use(middleware.Authorization)
		router.Use(middleware.LogRequest)

		// Handler
		router.Get("/get", GetEmployeeById)
		router.Get("/", GetEmployees)
	})


	// These are global calls
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/logs", func(router chi.Router) {
		// Middleware
		router.Use(middleware.LogRequest)

		router.Get("/", GetLogs)
	})

}