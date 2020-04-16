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

// User store interface expose the User db methods/operations
type User interface {
	All() ([]*schema.User, *errors.AppError)
	Create(req *schema.UserReq) (*schema.User, *errors.AppError)
	// GetByID(cityID uint) (*schema.City, *errors.AppError)
	// Update(city *schema.City, update *schema.City) (*schema.City, *errors.AppError)
	// Delete(cityID uint) *errors.AppError
}

// Incident store interface expose the Incident db methods/operations
type Incident interface {
	All() ([]*schema.Incident, *errors.AppError)
	Create(req *schema.IncidentReq) (*schema.Incident, *errors.AppError)
	GetByID(incidentID uint) (*schema.Incident, *errors.AppError)
	// Update(city *schema.City, update *schema.City) (*schema.City, *errors.AppError)
	// Delete(cityID uint) *errors.AppError
}
