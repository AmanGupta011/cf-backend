package controllers

import(
	"context"
	"encoding/json"
	"net/http"
	"time"
	"fmt"

	"cf-backend/auth"
	"cf-backend/db"
	"cf-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode the JSON body into the credentials struct
	var credentials auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println(credentials)

	// Initialize a variable to hold the user data
	var user models.User

	// Get a database connection
	client, err := db.DbConnection()
	userCollection := client.Database("cfstress").Collection("users")

	// Find the user by email
	err = userCollection.FindOne(context.TODO(), bson.M{"email": credentials.Email}).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode("No user found")
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		json.NewEncoder(w).Encode("Password is incorrect")
		return
	}

	// Generate a JWT token
	token, err := auth.GenerateJWT(credentials.Email)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	// Set the JWT token as a cookie
	http.SetCookie(w,
		&http.Cookie{
			Name:    "cfstressjwt",
			Value:   token,
			Expires: time.Now().Add(24 * 28 * time.Hour),
			Path:    "/",
		})
	
	// Respond with the user data
	json.NewEncoder(w).Encode(user)
}