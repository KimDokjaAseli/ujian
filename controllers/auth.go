package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gocroot/core"
	"github.com/gocroot/mgdb"
)

var db = mgdb.NewDatabase("mongodb://localhost:27017", "backend-ujian")

// User model

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// SignUp handles user registration
func SignUp(ctx *core.Context) {
	var user User
	if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, core.H{"error": "Invalid request"})
		return
	}
	// Insert user into MongoDB
	_, err := db.Collection("User").InsertOne(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, core.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusOK, core.H{"message": "User created"})
}

// Login handles user authentication
func Login(ctx *core.Context) {
	var req User
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.H{"error": "Invalid request"})
		return
	}
	// Find user in MongoDB
	var user User
	err := db.Collection("User").FindOne(mgdb.M{"username": req.Username}).Decode(&user)
	if err != nil || user.Password != req.Password {
		ctx.JSON(http.StatusUnauthorized, core.H{"error": "Invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, core.H{"message": "Login successful"})
}
