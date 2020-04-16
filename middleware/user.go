package middleware

import (
	"buzzme/pkg/errors"
	"buzzme/pkg/respond"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// UserRequired validates
func UserRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userIDStr := chi.URLParam(r, "userID")
		userID, er := strconv.Atoi(userIDStr)
		fmt.Println(er)
		fmt.Println("-------------------- Roshan --------------------", userID)
		if er != nil {
			respond.Fail(w, errors.BadRequest("invalid id").AddDebug(er))
			return
		}

		user, err := Store.User().GetByID(uint(userID))
		if err != nil {
			respond.Fail(w, err)
			return
		}

		ctx := ContextWrapAll(r.Context(), map[interface{}]interface{}{
			"userID": uint(userID),
			"user":   user,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
