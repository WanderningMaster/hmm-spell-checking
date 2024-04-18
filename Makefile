dev:
	@go run cmd/main.go
build:
	@go build -o bin/main cmd/main.go
run:
	@bin/main
clean_data:
	@go run ./data/clean.go
test:
	@go test -v ./... | grep -v "data" | grep -v "cmd"
