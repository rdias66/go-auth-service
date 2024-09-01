package routes

import (
	"go-auth-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/refresh-token", controllers.RefreshToken).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("POST")
}
