# Variables
BINARY_NAME=fullstack-app

.DEFAULT_GOAL := dev

# Commands
# ------------------------------------------------------------------------------------------
### Install all go dependencies ###
deps:
	go mod download && cd web && pnpm install

### Update all go dependencies ###
vet:
	go vet ./...

# ------------------------------------------------------------------------------------------
### Generate swagger docs for the api ###
swagger: deps
	swag init -g main.go -o swagger -d cmd/api,internal/api/handlers,internal/api/models,internal/api/models/dtos

### Run in development mode ###
dev: swagger
	cd web && pnpm dev & air && fg

# ------------------------------------------------------------------------------------------
### Build for all the platforms ###
build: swagger
	GOARCH=amd64 GOOS=darwin go build -o ./out/${BINARY_NAME}-darwin cmd/api/main.go
	GOARCH=amd64 GOOS=linux go build -o ./out/${BINARY_NAME}-linux cmd/api/main.go
	GOARCH=amd64 GOOS=windows go build -o ./out/${BINARY_NAME}-windows cmd/api/main.go

### Run binary for current platform ###
run: build
	./out/${BINARY_NAME}-darwin

### Clean all generated binaries ###
clean:
	go clean
	rm -rf ./out

# ------------------------------------------------------------------------------------------
### Test all go files ###
test:
	go test ./...

### Generate test coverage ###
coverage:
	go test ./... -coverprofile=coverage.out

# ------------------------------------------------------------------------------------------

