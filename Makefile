default: help

.PHONY: help 
help: ## print this help message: see https://stackoverflow.com/a/64996042
	@echo 'Usage:'
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-22s\033[0m %s\n", $$1, $$2}'

.PHONY: confirm 
confirm: ## confirm y or N
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: build
build: ## build the main app to ./bin/main
	go build -o ./bin/main ./cmd

.PHONY: run
run: ## run the cmd/api application
	go run ./cmd


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

.PHONY: audit
audit: vendor ## tidy and vendor dependencies and format, vet and test all code
	@echo 'Formatting code…'
	go fmt ./...
	@echo 'Vetting code…'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests…'
	go test -race -vet=off ./...

.PHONY: vendor
vendor: ## tidy and vendor dependencies
	@echo 'Tidying and verifying module dependencies…'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies…'
	go mod vendor