postgres:
	docker run -p 5432:5432 --name bank -e POSTGRES_USER=root -e POSTGRES_PASSWORD=8084689296 -d postgres:14-alpine
createdb:
	docker exec -it bank createdb --username=root --owner=root bank
dropdb:
	docker exec -it bank dropdb bank
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
migrateup:
	migrate -path db/migration -database "postgresql://root:Ng4iHN2A7es0hGsfrpQY@bank-db.c3pvcgxp5zez.ap-northeast-1.rds.amazonaws.com/bank_db" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:Ng4iHN2A7es0hGsfrpQY@bank-db.c3pvcgxp5zez.ap-northeast-1.rds.amazonaws.com/bank_db -verbose down
sqlcgen:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mockdb:
	mockgen -package mockdb  -destination db/mock/store.go  github.com/MohammadZeyaAhmad/Bank-App/db/sqlc Store
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine
.PHONY:
	postgres createdb dropdb migrateup migratedown sqlcgen test server mockdb proto redis new_migration