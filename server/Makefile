BINARY_NAME = server
AIR_BIN = $(HOME)/go/bin/air

all: build

build:
	@echo "Building the project..."
	go build -o $(BINARY_NAME) ./cmd

run: build
	@echo "Running the server..."
	./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)


fmt:
	@echo "Formatting Go code..."
	go fmt ./...

lint:
	@echo "Linting Go code..."
	golangci-lint run

help:
	@echo "Available commands:"
	@echo "  make build  - Build the project"
	@echo "  make run    - Build and run the server"
	@echo "  make clean  - Remove build artifacts"
	@echo "  make fmt    - Format Go code"
	@echo "  make lint   - Lint Go code"
	@echo "  make help   - Show this help message"