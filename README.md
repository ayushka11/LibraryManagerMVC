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
- `CREATE DATABASE dbname;` 
- `USE dbname;`

#### Running the server

- `migrate -path ./migrations -database "mysql://user:password@tcp(localhost:3306)/dbname" up`
- `go build -o mvc ./cmd/main.go`
- `./mvc`

## To run using script

- Make sure you have the above mentioned downloads installed.
- go to the project directory, and make the script executable by `chmod +x setup.sh`
- Run the script using `./setup.sh`

## Hosting

- Install apache2: `sudo apt install apache2`
- `sudo a2enmod proxy proxy_http`
- `cd /etc/apache2/sites-available`
- `sudo nano mvc.sdslabs.local.conf`
Add:

```xml
<VirtualHost *:80>
	ServerAdmin youremailid
	ProxyPreserveHost On
	ProxyPass / http://127.0.0.1:8080/
	ProxyPassReverse / http://127.0.0.1:8080/
	TransferLog /var/log/apache2/mvc_access.log
	ErrorLog /var/log/apache2/mvc_error.log
</VirtualHost>
```

- `sudo a2ensite mvc.sdslabs.local.conf`
- add `127.0.0.1	mvc.sdslabs.local` to `/etc/hosts`
- `sudo a2dissite 000-default.conf`
- `sudo apache2ctl configtest `
- `sudo systemctl restart apache2`
- `sudo systemctl status apache2`
