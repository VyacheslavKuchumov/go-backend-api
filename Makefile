build:
	@go build -o bin/backend_app cmd/main.go

test:
	@go test ./...

run: build
	@./bin/backend_app