include .env

migration:
	@echo "Generating migration file"
	atlas migrate diff --env gorm

up:
	@echo "Upgrading migration"
	atlas migrate apply \
	--url $(DATABASE_URL)

down:
	@echo "Downgrading migration"
	atlas migrate down \
	--url $(DATABASE_URL) \
	--dev-url $(CLEAN_DATABASE_URL)

hash:
	@echo "Recalculating sum"
	atlas migrate hash

dev:
	docker-compose up -d

dev--build:
	docker-compose up --build

prod:
	go run .