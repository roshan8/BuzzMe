package cities

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/weather-monster/api"
	"github.com/weather-monster/api/v1/forecasts"
	"github.com/weather-monster/middleware"
)

// Init initializes all the v1 routes
func Init(r chi.Router) {
	store = api.Store

	// ROUTE: {host}/v1/cities
	r.Method(http.MethodGet, "/", api.Handler())
	r.Method(http.MethodPost, "/", api.Handler(createCityHandler))
	r.With(middleware.CityRequired).
		Route("/{cityID:[0-9]+}", cityIDSubRoutes)
}

// ROUTE: {host}/v1/cities/:cityID/*
func cityIDSubRoutes(r chi.Router) {
	r.Method(http.MethodGet, "/", api.Handler(getCityHandler))
	r.Method(http.MethodPatch, "/", api.Handler(updateCityHandler))
	r.Method(http.MethodDelete, "/", api.Handler(deleteCityHandler))

	r.Method(http.MethodGet, "/temperature", api.Handler(getCityTemperatureHandler))
	r.Method(http.MethodGet, "/webhooks", api.Handler(getCityWebhookHandler))
	r.Method(http.MethodGet, "/forecast", api.Handler(forecasts.WeatherForecastHandler))
}
