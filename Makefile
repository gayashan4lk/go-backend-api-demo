build:
	@go build -o bin/go-backend-api-demo cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-backend-api-demo