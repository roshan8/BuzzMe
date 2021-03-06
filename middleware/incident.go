package middleware

import (
	"buzzme/pkg/errors"
	"buzzme/pkg/respond"
	"buzzme/store"
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Store holds new store connection
var Store *store.Conn

// Init ...
func Init(st *store.Conn) {
	Store = st
}

// IncidentRequired validates
func IncidentRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		incidentIDStr := chi.URLParam(r, "incidentID")
		incidentID, er := strconv.Atoi(incidentIDStr)

		if er != nil {
			respond.Fail(w, errors.BadRequest("invalid id").AddDebug(er))
			return
		}

		incident, err := Store.Incident().GetByID(uint(incidentID))
		if err != nil {
			respond.Fail(w, err)
			return
		}

		ctx := ContextWrapAll(r.Context(), map[interface{}]interface{}{
			"incidentID": uint(incidentID),
			"incident":   incident,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

// ContextWrapAll is used to set the following values in the
// passed context
func ContextWrapAll(ctx context.Context, x map[interface{}]interface{}) context.Context {
	for key, value := range x {
		ctx = context.WithValue(ctx, key, value)
	}

	return ctx
}
