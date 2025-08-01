package main

import (
	"net/http"
	"ujian/controllers"
)

func main() {
	http.HandleFunc("/signup", controllers.SignUp)
	http.HandleFunc("/login", controllers.Login)

	// Enable CORS middleware
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
