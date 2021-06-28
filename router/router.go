package router

import (
	"learn-go/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("api/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("api/deleteuser/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")
	router.HandleFunc("api/user/newuser", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("api/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")

	return router
}
