postgres:
	sudo docker run --name auth-service -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=4200 -d postgres
createdb:
	sudo docker exec -it auth-service createdb --username=root --owner=root auth-service
dropdb: 
	sudo docker exec -it auth-service dropdb auth-service
migrateup:
	migrate -path db/migration -database "postgres://root:4200@localhost:5432/auth-service?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://root:4200@localhost:5432/auth-service?sslmode=disable" -verbose down
migrateup1:
	migrate -path db/migration -database "postgres://root:4200@localhost:5432/auth-service?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path db/migration -database "postgres://root:4200@localhost:5432/auth-service?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test

