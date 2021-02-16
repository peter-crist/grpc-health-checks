APP_NAME ?= angry_server
PORT ?= 5000

# Genrate the api.pb.go
generate-proto:
	protoc -I ./proto angry.proto --go_out=plugins=grpc:./proto

# DOCKER TASKS
# Build the container
build-server: ## Build the container
	docker build -t $(APP_NAME) -f ./server/Dockerfile .

run-server: ## Run container on port configured in `config.env`
	docker run -i -t --rm -p=$(PORT):$(PORT) --name="$(APP_NAME)" $(APP_NAME)


serve: build-server run-server ## Run container on port configured in `config.env` (Alias to run)