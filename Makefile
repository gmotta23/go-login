include .env

migration:
	@echo "Generating migration file"
	atlas migrate diff --env gorm

migrate:
	atlas migrate apply --url $(DATABASE_URL)

rollback:
	@echo "Rolling back migration"
	atlas migrate hash
	atlas schema apply \
		--url $(DATABASE_URL) \
		--to "file://migrations" \
		--dev-url "docker://postgres/15?search_path=public" \
		--exclude "atlas_schema_revisions"
