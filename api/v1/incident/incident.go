package incident

import (
	"net/http"

	"buzzme/pkg/errors"
	"buzzme/pkg/respond"
	"buzzme/schema"
	"buzzme/utils"
)

// InitUsers fetches and unmarshal the user data from yaml config files
func getAllIncidentsHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	incidents, err := store.Incident().All()
	if err != nil {
		return err
	}

	respond.OK(w, incidents)
	return nil
}

// creates a new incident
func createIncidentHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	var input schema.IncidentReq

	if err := utils.Decode(r, &input); err != nil {
		return errors.BadRequest(err.Error()).AddDebug(err)
	}

	city, err := store.Incident().Create(&input)
	if err != nil {
		return err
	}

	respond.Created(w, city)
	return nil
}

// Get incident details by ID
func getIncidentHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	incident, _ := ctx.Value("incident").(*schema.Incident)

	respond.OK(w, incident)
	return nil
}
