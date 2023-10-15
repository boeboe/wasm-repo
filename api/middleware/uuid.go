package middleware

import (
	"context"
	"fmt"
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
					fmt.Printf("Error parsing UUID: %v\n", value)
					http.Error(w, "Invalid UUID format", http.StatusBadRequest)
				}
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
