postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=${PG_USER} -e POSTGRES_PASSWORD=${PG_PASSWORD} -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=${PG_USER} --owner=${PG_USER} simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

sqlc:
	sqlc generate

migrateup:
	migrate -path infra/db/migration -database "postgres://${PG_USER}:${PG_PASSWORD}@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path infra/db/migration -database "postgres://${PG_USER}:${PG_PASSWORD}@localhost:5432/simple_bank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb sqlc migrateup migratedown test