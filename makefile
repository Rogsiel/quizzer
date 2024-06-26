postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root quizzer
dropdb:
	docker exec -it postgres12 dropdb quizzer

migrateup:
	migrate -path internal/database/migration -database "postgresql://root:secret@localhost:5432/quizzer?sslmode=disable" -verbose up
migratedown:
	migrate -path internal/database/migration -database "postgresql://root:secret@localhost:5432/quizzer?sslmode=disable" -verbose down

migrateup1:
	migrate -path internal/database/migration -database "postgresql://root:secret@localhost:5432/quizzer?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path internal/database/migration -database "postgresql://root:secret@localhost:5432/quizzer?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run ./cmd//quizzer/main.go

build:
	@go build -o bin/main cmd/quizzer/main.go 

.PHONY:postgres createdb dropdb migrateup migratedown sqlc
