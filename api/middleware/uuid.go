package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func UUIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()

		// List of known ID keys
		idKeys := []string{"pluginID", "releaseID"}

		for _, key := range idKeys {
			if value, ok := vars[key]; ok {
				if id, err := uuid.Parse(value); err == nil {
					ctx = context.WithValue(ctx, key, id)
				} else {
					// Handle invalid UUID format if necessary
					// For example, you can send an error response and return:
					// http.Error(w, "Invalid UUID format", http.StatusBadRequest)
					// return
				}
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
