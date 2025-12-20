build:
	@go build -o bin/backend_app cmd/main.go

test:
	@go test ./...

run:
	@go run cmd/main.go