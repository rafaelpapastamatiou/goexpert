createMigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert-17-sqlc" -verbose up

migrateDown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert-17-sqlc" -verbose down