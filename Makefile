CURRENT_DIR = ./

postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:		
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc_init:
	docker run --rm -v $(CURRENT_DIR):/src -w /src sqlc/sqlc init

sqlc_generate:
	docker run --rm -v $(CURRENT_DIR):/src -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb

.PHONY: migrateup migratedown

.PHONY: sqlc_init sqlc_generate

.PHONY: test