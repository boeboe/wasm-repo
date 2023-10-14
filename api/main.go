package main

import (
	"log"
	"net/http"

	"github.com/boeboe/wasm-repo/api/handlers"
	"github.com/boeboe/wasm-repo/api/middleware"
	"github.com/boeboe/wasm-repo/api/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	// Global middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.ErrorHandlingMiddleware)

	// Default base route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to this wasm-repo"))
	}).Methods("GET")

	// WASMPlugin routes
	r.HandleFunc("/plugins", handlers.ListAllPluginsHandler).Methods("GET")
	r.HandleFunc("/plugins", handlers.CreatePluginHandler).Methods("POST")
	r.HandleFunc("/plugins/{id}", handlers.GetPluginByIDHandler).Methods("GET")
	r.HandleFunc("/plugins/{id}", handlers.UpdatePluginHandler).Methods("PUT")
	r.HandleFunc("/plugins/{id}", handlers.DeletePluginHandler).Methods("DELETE")

	// WASMRelease routes
	r.HandleFunc("/plugins/{pluginID}/releases", handlers.ListAllReleasesForPluginHandler).Methods("GET")
	r.HandleFunc("/plugins/{pluginID}/releases", handlers.CreateReleaseForPluginHandler).Methods("POST")
	r.HandleFunc("/plugins/{pluginID}/releases/{releaseID}", handlers.GetReleaseByIDHandler).Methods("GET")
	r.HandleFunc("/plugins/{pluginID}/releases/{releaseID}", handlers.UpdateReleaseForPluginHandler).Methods("PUT")
	r.HandleFunc("/plugins/{pluginID}/releases/{releaseID}", handlers.DeleteReleaseForPluginHandler).Methods("DELETE")

	// Start the server
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
