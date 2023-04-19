run:
	go run ./cmd/api/

dup:
	docker-compose up --build -d

postgres:
	docker run --name listingsdb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:15-alpine

createdb:
	docker exec -it listingsdb createdb --username=listingsdb --owner=root listingsdb

mup:
	migrate -path migrations -database "postgres://listingsdb:postgres@localhost/listingsdb?sslmode=disable" -verbose up 

mdown:
	migrate -path migrations -database "postgres://listingsdb:postgres@localhost/listingsdb?sslmode=disable" -verbose down
