package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/boeboe/wasm-repo/api/errors"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("An error occurred: %v", err)

				switch e := err.(type) {
				case *errors.ValidationError:
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(ErrorResponse{Status: http.StatusBadRequest, Message: e.Error()})
				case *errors.JSONDecodingError:
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(ErrorResponse{Status: http.StatusBadRequest, Message: e.Error()})
				case *errors.DatabaseError:
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(ErrorResponse{Status: http.StatusInternalServerError, Message: e.Error()})
				default:
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error"})
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}
