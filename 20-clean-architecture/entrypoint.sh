#!/bin/bash
# wait a little bit for the database to be ready
echo "Waiting for the database to be ready..."
sleep 5

# run the migrations
echo "Running database migrations..."
migrate -path=internal/infra/sql/migrations -database "mysql://root:root@tcp(mysql:3306)/orders" -verbose up

# install dependencies
echo "Installing dependencies..."
go mod tidy

# start the application
echo "Starting the application..."
cd cmd/ordersystem
go run main.go wire_gen.go