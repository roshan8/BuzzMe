package store

import (
	"buzzme/pkg/errors"
	"buzzme/schema"
)

// Store global store interface - provides db intercae methods
// for diff entities
type Store interface {
	User() User
	Incident() Incident
	// TODO: escalation policy and others
}

// Cities store interface expose the city db methods
type User interface {
	All() ([]*schema.Users, *errors.AppError)
	// Create(req *schema.CityReq) (*schema.City, *errors.AppError)
	// GetByID(cityID uint) (*schema.City, *errors.AppError)
	// Update(city *schema.City, update *schema.City) (*schema.City, *errors.AppError)
	// Delete(cityID uint) *errors.AppError
}

type Incident interface {
	All() ([]*schema.Incident, *errors.AppError)
	Create(req *schema.IncidentReq) (*schema.Incident, *errors.AppError)
	// GetByID(cityID uint) (*schema.City, *errors.AppError)
	// Update(city *schema.City, update *schema.City) (*schema.City, *errors.AppError)
	// Delete(cityID uint) *errors.AppError
}
