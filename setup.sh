#!/bin/bash

command_exists() {
    command -v "$1" >/dev/null 2>&1
}

generate_jwt_secret_key() {
    openssl rand -base64 32
}

if command_exists go; then
    echo "Go is installed."
else
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

if command_exists migrate; then
    echo "golang-migrate is installed."
else
    echo "golang-migrate is not installed. Please install golang-migrate and try again."
    exit 1
fi

if command_exists mysql; then
    echo "MySQL is installed."
else
    echo "MySQL is not installed. Please install MySQL and try again."
    exit 1
fi

read -p "Enter MySQL username: " DB_USERNAME
read -sp "Enter MySQL password: " DB_PASSWORD
read -p "Enter MySQL host: " DB_HOST
echo
read -p "Enter MySQL database name: " DB_NAME
# read -p "Enter Server Port: " SERVER_PORT

read -p "Do you have a JWT secret key? (yes/no): " HAS_JWT_SECRET_KEY

if [[ "$HAS_JWT_SECRET_KEY" == "yes" ]]; then
    read -p "Enter JWT Secret Key: " JWT_SECRET_KEY
else
    JWT_SECRET_KEY=$(generate_jwt_secret_key)
    echo "Generated JWT Secret Key: $JWT_SECRET_KEY"
fi

cat <<EOF >db.yaml
  DB_USERNAME: ${DB_USERNAME}
  DB_PASSWORD: ${DB_PASSWORD}
  DB_HOST: ${DB_HOST}
  DB_NAME: ${DB_NAME}
  JWTSecretKey: ${JWT_SECRET_KEY}
EOF

echo "Configuration file db.yaml created."

export DB_USERNAME
export DB_PASSWORD
export DB_NAME

echo "Creating database if it doesn't exist..."
mysql -u "$DB_USERNAME" -p"$DB_PASSWORD" -h 127.0.0.1 -e "CREATE DATABASE IF NOT EXISTS $DB_NAME;"

echo "Running database migrations..."
migrate -path ./migrations -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp(localhost:3306)/$DB_NAME" up

TIDY_CMD="go mod tidy"
echo "Tidying up the dependencies..."
if $TIDY_CMD; then
    echo "Tidy successful."
else
    echo "Tidy failed."
    exit 1
fi

VENDOR_CMD="go mod vendor"
echo "Vendoring dependencies..."
if $VENDOR_CMD; then
    echo "Vendoring successful."
else
    echo "Vendoring failed."
    exit 1
fi

BUILD_CMD="go build -o mvc ./cmd/main.go"

if [ -f librarymanager ]; then
    echo "Removing old build..."
    rm librarymanager
fi

echo "Building the project..."
if $BUILD_CMD; then
    echo "Build successful."
else
    echo "Build failed."
    exit 1
fi

RUN_CMD="./mvc"
echo "Starting the application..."
if $RUN_CMD; then
    echo "Application started."
else
    echo "Failed to start the application."
    exit 1
fi
