POSTGRESQL_URL ?= postgres://postgres:123@localhost:5432/docks?sslmode=disable=

build:
	go build main.go

start:
	go run main.go

test:
	go clean -testcache
	go test ./... -v

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path ./infra/database/migrations up

