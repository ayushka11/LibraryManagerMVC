# LibraryManagerMVC

This project is a Library Management System using MVC architecture and mySQL written in go.

### Download

- GO
- mySQL
- golang-migrate

### To run 

- Clone the repository
- In the repo directory, run `go mod vendor` and `go mod tidy`

#### MySQL

- `mysql -u root -p` and enter password
- `CREATE DATABASE lms;` 
- `USE lms;`

#### Running the server

- `migrate -path ./migrations -database "mysql://user:password@tcp(localhost:3306)/library_management_system" up`
- `go build -o mvc ./cmd/main.go`
- `./mvc`