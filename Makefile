SHELL := /bin/bash

# Include variables from the .envrc file
include .envrc

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

run:
	go run ./cmd/api/ -db-dsn=${DSN}

dup:
	docker-compose up --build -d

## cremig: create migration; needs an ragument name=custom_migration_name
cremig:
	migrate create -seq -ext=.sql -dir=./migrations ${name}

mversion:
	migrate -path migrations -database ${DSN} version

mup:
	migrate -path migrations -database ${DSN} -verbose up 

mdown: confirm
	migrate -path migrations -database ${DSN} -verbose down

mdownone: confirm
	migrate -path migrations -database ${DSN} -verbose down 1

itdb:
	docker exec -it listingsdb psql -U listingsdb

.PHONY: help confirm run dup cremig mversion mup mdown mdownone itdb