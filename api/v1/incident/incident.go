package incident

import (
	"net/http"

	"buzzme/pkg/errors"
	"buzzme/pkg/respond"
)

// InitUsers fetches and unmarshal the user data from yaml config files
func getAllIncidentsHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {

	users, err := store.Incident().All()
	if err != nil {
		return err
	}

	respond.OK(w, users)
	return nil
}
