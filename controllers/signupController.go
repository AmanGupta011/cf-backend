package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cf-backend/db"
	"cf-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (app *Application) SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode the JSON body into the user struct
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	// Get a database connection
	client, err := db.DbConnection()
	userCollection := client.Database("cfstress").Collection("users")

	// Check if the user already exists	
	res := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email})
	if res.Err() == nil {
		json.NewEncoder(w).Encode("This email already exists")
		fmt.Println(err)
		return
	}

	// Hash the user's password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		json.NewEncoder(w).Encode("Password does not match")
		return
	}
	hashed := string(hash)
	user.Password = hashed

	// Insert the user into the database
	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	// Respond with success message
	json.NewEncoder(w).Encode("success")
}