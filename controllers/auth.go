package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gocroot/mgdb"
	"go.mongodb.org/mongo-driver/bson"
)

var db, _ = mgdb.MongoConnect(mgdb.DBInfo{
	DBString: "mongodb://localhost:27017",
	DBName:   "backend-ujian",
})

// User model

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// SignUp handles user registration
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Insert user into MongoDB
	_, err := db.Collection("User").InsertOne(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User created"}`))
}

// Login handles user authentication
func Login(w http.ResponseWriter, r *http.Request) {
	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Find user in MongoDB
	var user User
	err := db.Collection("User").FindOne(r.Context(), bson.M{"username": req.Username}).Decode(&user)
	if err != nil || user.Password != req.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Login successful"}`))
}
