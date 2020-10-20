package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Route defines the Mux
// router individual route
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Index is the handler for the path "/"
func CrawlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World GoCrawler Node\n")
}

// SpiderIndex is the handler for the path "/"
func SpiderIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World GoSupervisor Node\n")
}

// Logger is the middleware to
// log the incoming request
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL,
			"method": r.Method,
		}).Info("incoming request")

		next.ServeHTTP(w, r)
	})
}

// Router returns a mux router
func Router(routes []Route) *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.HandleFunc(
			route.Path,
			route.Handler,
		).Methods(route.Method)
	}

	router.Use(Logger)

	return router
}
