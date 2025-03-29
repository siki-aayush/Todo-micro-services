TODO_SERVICE_BINARY=todoServiceApp
LOGGER_BINARY=loggerServiceApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_todo_service build_logger build_frontend
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_todo_service:
	@echo "Building todo service binary..."
	cd ../todo-service && env GOOS=linux CGO_ENABLED=0 go build -o ${TODO_SERVICE_BINARY} ./cmd/api
	@echo "Done!"

build_logger:
	@echo "Building logger binary..."
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"

build_frontend:
	@echo "Building frontend..."
	cd ../todo-app && bun run build
	@echo "Done!"