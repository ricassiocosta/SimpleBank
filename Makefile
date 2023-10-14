postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASS} -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=${DB_USER} --owner=${DB_USER} simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

sqlc:
	sqlc generate

migrateup_ci:
	migrate -path infra/db/migration -database "${{ secrets.POSTGRES_CONN_URL }}" -verbose up

migrateup:
	migrate -path infra/db/migration -database "postgres://${DB_USER}:${DB_PASS}@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path infra/db/migration -database "postgres://${DB_USER}:${DB_PASS}@localhost:5432/simple_bank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb sqlc migrateup migratedown test