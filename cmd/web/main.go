package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/raunchers/workAutomation/internal/handlers"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8000"

var sessionManager *scs.SessionManager

func main() {

	// Initialize a new session manager and configure the session lifetime.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	mux := http.NewServeMux()

	// Create file server which serves files out of the ./templates directory.
	fileServer := http.FileServer(http.Dir("./templates"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/templates/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/templates", http.StripPrefix("/templates", fileServer))

	mux.HandleFunc("/", handlers.Home)

	mux.HandleFunc("/OfflineRTR", handlers.OfflineRTR)
	mux.HandleFunc("/OfflineASTC", handlers.OfflineASTC)
	mux.HandleFunc("/OfflineWAN2", handlers.OfflineWAN2)
	mux.HandleFunc("/OfflineWAN3", handlers.OfflineWAN3)
	mux.HandleFunc("/OfflineIPSEC", handlers.OfflineIPSEC)

	mux.HandleFunc("/PostOfflineRTR", handlers.PostOfflineRTR)
	mux.HandleFunc("/PostOfflineASTC", handlers.PostOfflineASTC)
	mux.HandleFunc("/PostOfflineWAN2", handlers.PostOfflineWAN2)
	mux.HandleFunc("/PostOfflineWAN3", handlers.PostOfflineWAN3)
	mux.HandleFunc("/PostOfflineIPSEC", handlers.PostOfflineIPSEC)

	mux.HandleFunc("/BouncingRTR", handlers.BouncingRTR)
	mux.HandleFunc("/BouncingASTC", handlers.BouncingASTC)
	mux.HandleFunc("/BouncingWAN2", handlers.BouncingWAN2)
	mux.HandleFunc("/BouncingWAN3", handlers.BouncingWAN3)
	mux.HandleFunc("/BouncingIPSEC", handlers.BouncingIPSEC)

	mux.HandleFunc("/PostBouncingRTR", handlers.PostBouncingRTR)
	mux.HandleFunc("/PostBouncingASTC", handlers.PostBouncingASTC)
	mux.HandleFunc("/PostBouncingWAN2", handlers.PostBouncingWAN2)
	mux.HandleFunc("/PostBouncingWAN3", handlers.PostBouncingWAN3)
	mux.HandleFunc("/PostBouncingIPSEC", handlers.PostBouncingIPSEC)

	log.Println("Starting server on:" + portNumber)
	log.Fatal(http.ListenAndServe(portNumber, sessionManager.LoadAndSave(mux)))
}
