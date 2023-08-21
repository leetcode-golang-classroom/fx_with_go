.Phony:

build:
	@go build -o bin/fx-go cmd/main.go

run: build
	@./bin/fx-go
