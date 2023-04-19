run:
	go run ./cmd/api/

dup:
	docker-compose up --build -d

mup:
	migrate -path migrations -database "postgres://listingsdb:postgres@localhost/listingsdb?sslmode=disable" -verbose up 

mdown:
	migrate -path migrations -database "postgres://listingsdb:postgres@localhost/listingsdb?sslmode=disable" -verbose down

mversion:
	migrate -path migrations -database "postgres://listingsdb:postgres@localhost/listingsdb?sslmode=disable" version

itdb:
	docker exec -it listingsdb psql -U listingsdb