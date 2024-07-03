package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type DBInfo struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}

type Claims struct {
	Username string 
	UserId   int    
	IsAdmin  bool   
	jwt.RegisteredClaims
}

type User struct {
	UserId   int
	UserName string
	IsAdmin  bool
	Password  string
	AdminRequestStatus string
}

type RequestUser struct {
	UserName string
	Password string
	ConfirmPassword string
}

type Book struct {
	BookId        int
	Title         string
	Author        string
	Available     int
	Quantity      int
}

type PgMessage struct {
	Message interface{}
}

type FileName struct {
	AdminHome           string
	UserHome            string
	Login               string
	PageNotFound        string
	UnauthorizedAccess  string
	InternalServerError string
	Signup              string
	AddBook             string
}

