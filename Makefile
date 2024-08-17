postgres:
	docker run --name authentication -p 5435:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine

createdb:
	docker exec -it authentication createdb --username=root --owner=root auth

dropdb:
	docker exec -it authentication dropdb spices-api

migrateup:
	migrate -path migrations -database "postgresql://root:secret@localhost:5435/auth?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://root:secret@localhost:5435/auth?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/server/main.go

.PHONY: postgres createdb dropdb migrateup migratedown test sqlc server