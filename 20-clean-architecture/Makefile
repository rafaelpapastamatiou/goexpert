protogen:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/*

createMigration:
	migrate create -ext=sql -dir=internal/infra/sql/migrations -seq init

migrate:
	migrate -path=internal/infra/sql/migrations -database "mysql://root:root@tcp(mysql:3306)/orders" -verbose up

migrateDown:
	migrate -path=internal/infra/sql/migrations -database "mysql://root:root@tcp(mysql:3306)/orders" -verbose down