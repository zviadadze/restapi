build:
	@go build -o bin/userver ./cmd/userver

run: build
	@./bin/userver