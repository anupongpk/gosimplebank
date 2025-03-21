startdocker:
	docker start postgres12

stopdocker:
	docker stop postgres12

postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=localhost -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=localhost --owner=localhost simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration/ -database "postgresql://localhost:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://localhost:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY:startdocker stopdocker postgres createdb dropdb migrateup migratedown sqlc test