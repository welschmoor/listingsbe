SHELL := /bin/bash

include .envrc


# ==================================================================================== # 
# HELPERS
# ==================================================================================== #

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: help confirm run dup cremig mversion mup mdown mdownone itdb audit vendor


# ==================================================================================== # 
# DEVELOPMENT
# ==================================================================================== #

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


# ==================================================================================== # 
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code; needs staticcheck installed
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
vendor:
	@echo 'Tidying and verifying module dependencies...' 
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor