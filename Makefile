APP_NAME ?= angry_server
CLIENT_NAME ?= chill_client
PORT ?= 5000

# Genrate the api.pb.go
generate-proto:
	protoc -I ./proto angry.proto --go_out=plugins=grpc:./proto

# DOCKER TASKS
# Build the container
build-server: ## Build the container
	docker build -t $(APP_NAME) -f ./server/Dockerfile .

run-server: ## Run container on port
	docker run -i -t --rm -p=$(PORT):$(PORT) --name="$(APP_NAME)" $(APP_NAME)

serve: build-server run-server ## Run container on port configured in `config.env` (Alias to run)

build-client:
	docker build -t $(CLIENT_NAME) -f ./client/Dockerfile .

run-client:
	docker run -it --rm --network="host" $(CLIENT_NAME)

client-up: build-client run-client