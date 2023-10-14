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
	repo := models.ConnectDatabase()

	r := mux.NewRouter()

	// Global middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.ErrorHandlingMiddleware)
	r.Use(middleware.UUIDMiddleware)

	// Initialize handler structs with the repository
	pluginHandler := &handlers.WASMPluginHandler{Repo: repo}
	releaseHandler := &handlers.WASMReleaseHandler{Repo: repo}

	// Default base route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to this wasm-repo"))
	}).Methods("GET")

	// WASMPlugin routes
	r.HandleFunc("/plugins", pluginHandler.ListAllPluginsHandler).Methods("GET")
	r.HandleFunc("/plugins", pluginHandler.CreatePluginHandler).Methods("POST")
	r.HandleFunc("/plugins/{pluginID}", pluginHandler.GetPluginByIDHandler).Methods("GET")
	r.HandleFunc("/plugins/{pluginID}", pluginHandler.UpdatePluginHandler).Methods("PUT")
	r.HandleFunc("/plugins/{pluginID}", pluginHandler.DeletePluginHandler).Methods("DELETE")

	// WASMRelease routes
	r.HandleFunc("/plugins/{pluginID}/releases", releaseHandler.ListAllReleasesForPluginHandler).Methods("GET")
	r.HandleFunc("/plugins/{pluginID}/releases", releaseHandler.CreateReleaseForPluginHandler).Methods("POST")
	r.HandleFunc("/plugins/{pluginID}/releases/{releaseID}", releaseHandler.GetReleaseByIDHandler).Methods("GET")
	r.HandleFunc("/plugins/{pluginID}/releases/{releaseID}", releaseHandler.UpdateReleaseForPluginHandler).Methods("PUT")
	r.HandleFunc("/plugins/{pluginID}/releases/{releaseID}", releaseHandler.DeleteReleaseForPluginHandler).Methods("DELETE")

	// Start the server
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
