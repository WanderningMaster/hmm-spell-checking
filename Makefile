dev:
	@air
build:
	@go build -o bin/main main.go
run:
	@bin/main
clean_data:
	@go run ./data/clean.go
test:
	@go test -v ./... | grep -v "data" | grep -v "cmd"
measure:
	@go run ./measures
