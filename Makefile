postgres:
	docker run --name postgres12  --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank

createmigration:
	migrate -help
	migrate create -ext sql -dir db/migration -seq [ADDNAME]

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/fxmbx/go-simple-bank/db/sqlc Store
builddocker:
	docker build -t simplebank:latest .
dockerrun:
	docker run --name simplebank --network bank-network  -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:root@postgres12:5432/simple_bank?sslmode=disable" simplebank:latest 

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc server mock builddocker dockerrun