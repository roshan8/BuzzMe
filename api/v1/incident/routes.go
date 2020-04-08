package incident

import (
	"net/http"

	"buzzme/api"
	"buzzme/middleware"
	appstore "buzzme/store"

	"github.com/go-chi/chi"
)

// store holds shared store conn from the api
var store *appstore.Conn

// Init initializes all the v1 routes
func Init(r chi.Router) {

	store = api.Store

	r.Method(http.MethodGet, "/", api.Handler(getAllIncidentsHandler))
	r.Method(http.MethodPost, "/", api.Handler(createIncidentHandler))
	r.With(middleware.IncidentRequired).
		Route("/{incidentID:[0-9]+}", incidentIDSubRoutes)
}

// ROUTE: {host}/v1/incident/:incidentID/*
func incidentIDSubRoutes(r chi.Router) {
	r.Method(http.MethodGet, "/", api.Handler(getIncidentHandler))
	// r.Method(http.MethodPatch, "/", api.Handler(updateCityHandler))
	// r.Method(http.MethodDelete, "/", api.Handler(deleteCityHandler))

	// r.Method(http.MethodGet, "/temperature", api.Handler(getCityTemperatureHandler))
	// r.Method(http.MethodGet, "/webhooks", api.Handler(getCityWebhookHandler))
	// r.Method(http.MethodGet, "/forecast", api.Handler(forecasts.WeatherForecastHandler))
}
