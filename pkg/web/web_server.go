package web

import (
	"github.com/gorilla/mux"
	"cf-backend/controllers"
	"cf-backend/pkg/store"
)

// CreateWebServer initializes the application and sets up the router with API endpoints
func CreateWebServer(counter int, ticketStore store.TicketStore) (*controllers.Application, *mux.Router) {
	// Buffer capacity for the application's boolean channel
	const bufferCapacity = 50

	// Initialize the application with the counter, ticket store, and buffered channel
	app := &controllers.Application{
		Counter:     counter,
		TicketStore: ticketStore,
		Channel:     make(chan bool, bufferCapacity),
	}

	// Initialize a new Gorilla Mux router
	r := mux.NewRouter()

	// Define API endpoints and their corresponding handlers
	r.HandleFunc("/api/signup", app.SignupHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/login", app.LoginHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/user", app.UserHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/logout", app.LogoutHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/contact", app.ContactController).Methods("POST", "OPTIONS")

	// // Define additional endpoints for test and status functionalities
	r.HandleFunc("/api/test/{contestID}/{problemIndex}", app.TestHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/status/{ticketID}", app.StatusHandler).Methods("GET")

	// Return the initialized application and router
	return app, r
}
