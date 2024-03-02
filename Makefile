SHELL=bash
.SHELLFLAGS=-ec -o pipefail

include .env
include .local.env

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build-images
build-images: ## Build docker-compose images   
build-images: 
	docker-compose build

.PHONY: run-images
run-images: ## Run docker-compose services   
run-images: 
	docker-compose up