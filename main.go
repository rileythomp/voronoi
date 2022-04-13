package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/rileythomp14/voronoi/src/handlers"
)

func main() {
	logger := handlers.NewLogger()
	s := http.Server{
		Addr:              ":8090",
		Handler:           createRouter(),
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      time.Minute,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}
	logger.Infof("Starting http server on port%s", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic("Error starting server: " + err.Error())
	}
}

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := handlers.GetRoutes()
	for _, r := range routes {
		router.Methods(r.Method).Path(r.Pattern).Name(r.Name).Handler(r.HandlerFunc)
	}
	return router
}
