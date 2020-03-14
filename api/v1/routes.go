package v1

import (
	//TODO: Need to import only v1 directly and use all the methods
	"github.com/go-chi/chi"
)

// Routes registered routes
func Routes(r chi.Router) {
	// r.Method(http.MethodGet, "/", v1.Handler(api.IndexHandeler))
	// r.Get("/top", api.HealthHandeler)
	r.Route("/v1", Init)
}

// Init initializes all the v1 routes
func Init(r chi.Router) {
	r.Route("/users", users.getAllUsersHandler)
}
