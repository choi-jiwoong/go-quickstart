.PHONY: build test run clean tidy coverage test-unit test-integration docker-build docker-run docker-compose-up docker-compose-down

build:
	go build -o bin/server ./cmd/server

test:
	go test ./...

test-unit:
	go test ./internal/repository -v
	go test ./internal/api -run "^Test[^Integration]" -v

test-integration:
	go test ./internal/api -run ".*Integration" -v

run:
	go run ./cmd/server

clean:
	rm -rf bin/

tidy:
	go mod tidy

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

docker-build:
	docker build -t go-quickstart .

docker-run:
	docker run -p 8080:8080 go-quickstart

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down