.PHONY: all build run watch clean swag swag-init swag-fmt

# Default: 
all: build

# Build
build:
	@go build -o ./tmp/main main.go

# Run 
run:
	@go run main.go

# Live Reload
dev:
	@AIR_PATH=$$(go env GOPATH)/bin/air; \
	if [ -f "$$AIR_PATH" ] || command -v air > /dev/null; then \
		if [ -f "$$AIR_PATH" ]; then \
			$$AIR_PATH; \
		else \
			air; \
		fi; \
		echo "Watching..."; \
	else \
		read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" = "y" ] || [ "$$choice" = "Y" ]; then \
			go install github.com/air-verse/air@latest; \
			$(go env GOPATH)/bin/air; \
			echo "Watching..."; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi

# Swagger documentation
swag: swag-init swag-fmt

# Initialize Swagger documentation
swag-init:
	@swag init -g main.go -o ./docs

# Format Swagger documentation
swag-fmt:
	@swag fmt

# Clean
clean:
	@rm -rf ./tmp
	@echo "Cleaned build artifacts."
