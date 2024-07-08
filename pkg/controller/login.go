package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ayushka11/LibraryManagerMVC/pkg/models"
	"github.com/ayushka11/LibraryManagerMVC/pkg/types"
	"github.com/ayushka11/LibraryManagerMVC/pkg/views"
	"github.com/golang-jwt/jwt/v5"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	files := views.ViewFileNames()
	t := views.Render(files.Login)
	t.Execute(writer, nil)
}

func LoginUser(writer http.ResponseWriter, request *http.Request) {
	var RequestUser types.RequestUser

	RequestUser.UserName = request.FormValue("username")
	RequestUser.Password = request.FormValue("password")
	
	db, err := models.Connection()
	if err != nil {
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	userExists, user, err := models.UserExists(db, RequestUser.UserName)
	if err != nil {
		fmt.Println("Error in UserExists")
		log.Println(err)
		http.Redirect(writer, request, "/500", http.StatusSeeOther)
		return
	}

	if !userExists {
		http.Redirect(writer, request, "/login?message=user does not exist", http.StatusSeeOther)
		return
	} else {
		if !models.CheckPasswordHash(RequestUser.Password, user.Password) {
			http.Redirect(writer, request, "/login?message=incorrect password", http.StatusSeeOther)
			return
		} else {
			expirationTime := time.Now().Add(1*time.Hour)

			claims := &types.Claims{
				Username: user.UserName,
				UserId: user.UserId,
				IsAdmin: user.IsAdmin,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(expirationTime),
				},
			}

			key, err := models.GetJWTSecretKey()
			jwtKey := []byte(key)
			if err != nil {
				fmt.Printf("Error getting JWT secret key: %v\n", err)
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
				return
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			tokenStr, err := token.SignedString(jwtKey)
			if err != nil {
				fmt.Printf("Error signing JWT token: %v\n", err)
				http.Redirect(writer, request, "/500", http.StatusSeeOther)
				return
			}

			http.SetCookie(writer, &http.Cookie{
				Name: "token",
				Value: tokenStr,
				Expires: expirationTime,
			})

			if (user.IsAdmin) {
				http.Redirect(writer, request, "/admin/adminHome", http.StatusSeeOther)
			} else {
				http.Redirect(writer, request, "/user/userHome", http.StatusSeeOther)
			}
		}
	}
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	expirationTime := time.Now().Add(5 * time.Second)
	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Expires: expirationTime,
	})
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}