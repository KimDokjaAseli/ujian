package main

import (
	"github.com/gocroot/core"
	"github.com/gocroot/mgdb"
	"ujian/controllers"
	"net/http"
)

func main() {
	app := core.New()

	// Enable CORS for all routes
	app.Use(func(ctx *core.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if ctx.Request.Method == "OPTIONS" {
			ctx.Writer.WriteHeader(http.StatusOK)
			return
		}
		ctx.Next()
	})

	// Auth routes
	app.Post("/signup", controllers.SignUp)
	app.Post("/login", controllers.Login)

	app.Run(":8080") // Listen on port 8080
}
