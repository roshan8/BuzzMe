package v1

import (
	"buzzme/api"
	v1 "buzzme/api"
	user "buzzme/api/v1/users"
	"net/http"

	"github.com/go-chi/chi"
)

// Routes registered routes
func Routes(r chi.Router) {
	r.Route("/api/v1", Init)
	r.Method(http.MethodGet, "/", v1.Handler(api.IndexHandeler))
	// r.Get("/top", api.HealthHandeler)
}

// Init initializes all the v1 routes
func Init(r chi.Router) {
	r.Route("/users", user.Init)
	// TODO: remaining routes for escalation policy and others
}
