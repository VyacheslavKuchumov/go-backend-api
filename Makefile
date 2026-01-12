build:
	@go build -o bin/backend_app cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/backend_app