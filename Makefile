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
### Generate swagger docs for the app ###
swagger: deps
	swag init -g main.go -o api/openapi -d cmd/app,internal/app/handlers,internal/app/models,internal/app/dtos

### Run in development mode ###
dev: swagger
	cd web && pnpm dev & air && fg

# ------------------------------------------------------------------------------------------
### Build for all the platforms ###
build: swagger
	GOARCH=amd64 GOOS=darwin go build -o ./dist/${BINARY_NAME}-darwin cmd/app/main.go
	GOARCH=amd64 GOOS=linux go build -o ./dist/${BINARY_NAME}-linux cmd/app/main.go
	GOARCH=amd64 GOOS=windows go build -o ./dist/${BINARY_NAME}-windows cmd/app/main.go

### Run binary for current platform ###
run: build
	./dist/${BINARY_NAME}-darwin

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
### Database Related
migration_create:
	@[ -n "$(name)" ] || (echo "Error: Migration name is required. Usage: make migration_create name=<migration_name>" && exit 1)
	goose create $(name) sql --dir db/migrations

migration_up:
	goose up --dir db/migrations --env .env

migration_down:
	goose down --dir db/migrations --env .env
