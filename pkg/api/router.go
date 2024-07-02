package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/ayushka11/LibraryManagerMVC/pkg/middleware"
	"github.com/ayushka11/LibraryManagerMVC/pkg/controller"
)

func Run() {
	fmt.Println("Server Started at post http://localhost:8000/")

	router := mux.NewRouter()
	router.Use(middleware.TokenMiddleware)

	userRouter := router.PathPrefix("/user").Subrouter()
	adminRouter := router.PathPrefix("/admin").Subrouter()

	userRouter.Use(middleware.RoleMiddleware(false))
	adminRouter.Use(middleware.RoleMiddleware(true))

	router.HandleFunc("/signup", controller.Signup).Methods("GET")
	router.HandleFunc("/signup", controller.AddUser).Methods("POST")
	
	// router.HandleFunc("/login", controller.Login).Methods("GET")
	// router.HandleFunc("/login", controller.LoginUser).Methods("POST")

	// router.HandleFunc("/403", controller.Unauthorized).Methods("GET")
	// router.HandleFunc("/500", controller.InternalServerError).Methods("GET")

	http.ListenAndServe(":8000", router)
}

