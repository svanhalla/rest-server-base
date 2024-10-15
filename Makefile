# Variabler
PROJECT_NAME = $(notdir $(CURDIR))
OPENAPI_SPEC := api/api.yaml  # Din OpenAPI-specifikation
GENERATED_TYPES := api/models.gen.go
GENERATED_SERVER := api/server.gen.go
BINARY := base-rest-server
ZIP_FILE = dist/$(PROJECT_NAME).zip

# Verktyg
OAPI_CODEGEN := oapi-codegen

# Hämta version från git tag (om tillämpligt)
VERSION := $(shell git describe --tags --always --dirty)
BUILD_DATE := $(shell date +%Y-%m-%d)
USER := $(shell whoami)

DOCKERFILE := Dockerfile

# Kommandon

.PHONY: all build clean generate

# Standardmål
all: generate build

# Skapa katalogen "internal" om den inte finns
internal:
	@mkdir -p internal

# Generera Go-kod från OpenAPI-specifikationen
generate: internal $(OPENAPI_SPEC)
	@echo "Generating Go code from OpenAPI spec..."
	$(OAPI_CODEGEN) -generate types -o $(GENERATED_TYPES) --package=api $(OPENAPI_SPEC)
	$(OAPI_CODEGEN) -generate server -o $(GENERATED_SERVER) --package=api $(OPENAPI_SPEC)

# Bygg projektet
build: generate format
	@echo "Building Go binary..."
	go build -ldflags "-X main.version=$(VERSION) -X main.buildDate=$(BUILD_DATE) -X main.user=$(USER)" -o bin/$(BINARY) ./cmd/

# Rensa genererade filer
clean:
	@echo "Cleaning up generated files..."
	rm -f $(GENERATED_TYPES) $(GENERATED_SERVER) $(BINARY)

format:
	goimports -w internal/*.go
	goimports -w cmd/*.go

zip:
	@mkdir -p dist
	@echo "Zipping all project files from root directory into dist/..."
	zip -r $(ZIP_FILE) . -x "dist/*" "*.git/*" "*.DS_Store" ".idea/*"
	@echo "Files zipped into $(ZIP_FILE)"

# Docker build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(PROJECT_NAME) -f $(DOCKERFILE) .

# Clean build artifacts
clean-docker:
	@echo "Cleaning up..."
	rm -f api/models.gen.go api/server.gen.go $(PROJECT_NAME)

# Lint the application
.PHONY: lint
lint: format
	@golangci-lint run



