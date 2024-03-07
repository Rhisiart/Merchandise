build:
	@go build -o bin/merchandise

run: build
	@./bin/merchandise

test:
	@go test -v ./..