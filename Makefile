postgres:
	docker run -p 5432:5432 --name bank -e POSTGRES_USER=root -e POSTGRES_PASSWORD=8084689296 -d postgres:12-alpine
createdb:
	docker exec -it bank createdb --username=root --owner=root bank
dropdb:
	docker exec -it bank dropdb bank
migrateup:
	migrate -path db/migration -database "postgresql://root:8084689296@localhost:5432/bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:8084689296@localhost:5432/bank?sslmode=disable" -verbose down
sqlcgen:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...
.PHONY:
	postgres createdb dropdb migrateup migratedown sqlcgen test 