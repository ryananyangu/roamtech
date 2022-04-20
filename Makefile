VERSION=`git rev-parse HEAD`
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
APPLICATION=scrapper
PROJECT=romtech


.PHONY: help
help: ## - Show help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: build
build:	## - Build the smallest and secured golang docker image based on scratch (ARGS ARE APPLICATION SPECIFIC)
	@printf "\033[32m\xE2\x9c\x93 $(VERSION) | Build the smallest and secured golang docker image based on scratch\n\033[0m"
	@export DOCKER_CONTENT_TRUST=1 && docker build -f Dockerfile  -t "$(PROJECT)/"$(APPLICATION):"$(VERSION)" .

.PHONY: build-no-cache
build-no-cache:	## - Build the smallest and secured golang docker image based on scratch with no cache
	@printf "\033[32m\xE2\x9c\x93 $(VERSION) | Build the smallest and secured golang docker image based on scratch\n\033[0m"
	@export DOCKER_CONTENT_TRUST=1 && docker build --no-cache -f Dockerfile -t $(PROJECT)/$(APPLICATION):$(VERSION) .

.PHONY: ls
ls: ## - Listing images of the application with versions
	@printf "\033[32m\xE2\x9c\x93 $(VERSION) | Listing images of the application with versions !\n\033[0m"
	@docker image ls gcr.io/$(PROJECT)/$(APPLICATION)

.PHONY: run
run:	## - Run the smallest and secured golang docker image based on scratch
	@printf "\033[32m\xE2\x9c\x93 $(VERSION) | Running image outside environment\n\033[0m"
	@docker run -p 8080:8080 "$(PROJECT)/$(APPLICATION):$(VERSION)"

.PHONY: push-to-gcp
push-to-gcp:	## - Push docker image to gcr.io container registry
	# @gcloud auth application-default login
	# @gcloud auth configure-docker
	@printf "\033[32m\xE2\x9c\x93 $(VERSION) | Pushing image to GCP !\n\033[0m"
	@docker push $(PROJECT)/$(APPLICATION):$(VERSION)

.PHONY: deploy-to-gcp
deploy-to-gcp: ## - Deploying version to GCLOUD RUN
	@printf "\033[32m\xE2\x9c\x93 $(VERSION) | Deploying version to GCLOUD RUN !\n\033[0m"
	@gcloud run deploy $(APPLICATION) --platform managed --revision-suffix=$(VERSION) --region us-central1 --image $(PROJECT)/$(APPLICATION):$(VERSION)

.PHONY: run-tests
run-tests: ## - Running tests on environment
	@printf "\033[32m\xE2\x9c\x93 $(VERSION) | Running tests on environment !\n\033[0m"
	@go test -v -race -coverprofile=cover.out -covermode=atomic ./...

.PHONY: test-coverage
run-coverage : ## - Generating the test coverage report
	@printf "\033[32m\xE2\x9c\x93 $(VERSION) | coverage report html cover.html !\n\033[0m"
	@go tool cover -html=cover.out -o cover.html

.PHONY: run-env
run-env: ## - Running application on local env
	@printf "\033[32m\xE2\x9c\x93 $(VERSION) | Running application on local env !\n\033[0m"
	@docker-compose up -d