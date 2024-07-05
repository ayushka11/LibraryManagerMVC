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
	
	router.HandleFunc("/login", controller.Login).Methods("GET")
	router.HandleFunc("/login", controller.LoginUser).Methods("POST")

	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	userRouter.HandleFunc("/userHome", controller.UserHome).Methods("GET")
	userRouter.HandleFunc("/availableBooks", controller.AvailableBooks).Methods("GET")
	userRouter.HandleFunc("/requestCheckout", controller.CheckoutBook).Methods("POST")
	userRouter.HandleFunc("/requestAdmin", controller.RequestAdmin).Methods("GET")
	userRouter.HandleFunc("/history", controller.ViewHistory).Methods("GET")

	adminRouter.HandleFunc("/adminHome", controller.AdminHome).Methods("GET")
	adminRouter.HandleFunc("/addBook", controller.AddBookPage).Methods("GET")
	adminRouter.HandleFunc("/addBook", controller.AddBook).Methods("POST")
	adminRouter.HandleFunc("/viewBooks", controller.ViewBooks).Methods("GET")
	adminRouter.HandleFunc("/deleteBook", controller.DeleteBook).Methods("POST")
	adminRouter.HandleFunc("/bookRequests", controller.ViewBookRequests).Methods("GET")	
	adminRouter.HandleFunc("/approveRequest", controller.ApproveRequest).Methods("POST")
	adminRouter.HandleFunc("/declineRequest", controller.DeclineRequest).Methods("POST")
	adminRouter.HandleFunc("/adminRequests", controller.ViewAdminRequests).Methods("GET")
	adminRouter.HandleFunc("/approveAdminRequest", controller.ApproveAdminRequest).Methods("POST")
	adminRouter.HandleFunc("/rejectAdminRequest", controller.RejectAdminRequest).Methods("POST")																													

	// router.HandleFunc("/403", controller.Unauthorized).Methods("GET")
	// router.HandleFunc("/500", controller.InternalServerError).Methods("GET")

	http.ListenAndServe(":8000", router)
}

