SHELL := /bin/bash
DSN :="postgres://listingsdb:postgres@localhost/listingsdb?sslmode=disable"

run:
	go run ./cmd/api/

dup:
	docker-compose up --build -d

mversion:
	migrate -path migrations -database ${DSN} version

mup:
	migrate -path migrations -database ${DSN} -verbose up 

mdown:
	migrate -path migrations -database ${DSN} -verbose down

mdownone:
	migrate -path migrations -database ${DSN} -verbose down 1

itdb:
	docker exec -it listingsdb psql -U listingsdb