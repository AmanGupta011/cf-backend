package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"cf-backend/auth"
	"cf-backend/db"
	"cf-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (app *Application) UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Retrieve the JWT from the request cookies
	cookie, err := r.Cookie("cfstressjwt")
	if err != nil {
		json.NewEncoder(w).Encode("unauthenticated")
		return
	}
	tokenString := cookie.Value

	// Validate the JWT token
	err, claims := auth.ValidateToken(tokenString)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	// Retrieve user information from the database
	var user models.User
	client, err := db.DbConnection()
	userCollection := client.Database("cfstress").Collection("users")
	err = userCollection.FindOne(context.TODO(), bson.M{"email": claims.Email}).Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	// Respond with the user information
	json.NewEncoder(w).Encode(user)
}