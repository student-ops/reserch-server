BROKER_BINARY=brokerApp

refresh:
	@echo "refreshing"
up: database_up
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

database_up:
	@echo "Starting database Docker images..."
	docker-compose -f ./database/docker-compose.yml up -d

database_build:
	docker-compose -f ./database/docker-compose.yml down
	@echo "Building (when required) and starting docker images..."
	docker-compose -f ./database/docker-compose.yml up --build -d
	@echo "database iamges built and strted!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: database_build
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"