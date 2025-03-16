.PHONY: postgres createdb dropdb migrateup migratedown sqlc sqlcwin test

help:
	@echo "postgres: run postgres docker container"
	@echo "createdb: create simple_bank database"
	@echo "dropdb: drop simple_bank database"
	@echo "migrateup: run database migration"
	@echo "migratedown: rollback database migration"
	@echo "sqlc: generate sqlc code"
	@echo "sqlcwin: generate sqlc code for windows"
	@echo "test: run all tests"

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:latest

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank --username=root 

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

sqlcwin:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover	./...