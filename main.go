package main

import(
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
    "cf-backend/pkg/store/mongodb"
	"cf-backend/pkg/web"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ticketStore, counter, err := mongodb.NewMongoStore()
	if err != nil {
		fmt.Println(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}


	app, r := web.CreateWebServer(counter, ticketStore)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	cred := handlers.AllowCredentials()
	fmt.Println(ticketStore)
    fmt.Println(counter)
    fmt.Println(app.Counter)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins, cred)(r)))
	// gorilla/handlers is a collection of handlers (aka "HTTP middleware") for use with Go's net/http package
}

