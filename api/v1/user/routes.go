package user

import (
	"net/http"

	"buzzme/api"
	appstore "buzzme/store"

	"github.com/go-chi/chi"
)

// store holds shared store conn from the api
var store *appstore.Conn

// Init initializes all the v1 routes
func Init(r chi.Router) {

	store = api.Store

	// ROUTE: {host}/v1/cities
	r.Method(http.MethodGet, "/", api.Handler(getAllUsersHandler))
	r.Method(http.MethodPost, "/", api.Handler(createUserHandler))
	// r.With(middleware.CityRequired).
	// 	Route("/{cityID:[0-9]+}", cityIDSubRoutes)
}
