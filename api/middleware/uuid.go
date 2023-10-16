package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/boeboe/wasm-repo/api/errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type contextKey string

const (
	PluginIDKey  contextKey = "pluginID"
	ReleaseIDKey contextKey = "releaseID"
)

func UUIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()

		// List of known ID keys
		idKeys := []contextKey{PluginIDKey, ReleaseIDKey}

		for _, key := range idKeys {
			if value, ok := vars[string(key)]; ok {
				if id, err := uuid.Parse(value); err == nil {
					ctx = context.WithValue(ctx, key, id)
				} else {
					panic(&errors.ValidationError{Field: string(key), Message: fmt.Sprintf("error parsing uuid: %v", err.Error())})
				}
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
