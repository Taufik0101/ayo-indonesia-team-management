generate-migration: ## Generate migration file example: generate-migration name=migration_name
	@[ "${name}" ] || ( echo "migration name not set"; exit 1 )
	migrate create -ext sql -dir database/migrations -seq -digits 5 $(name)

migrate-up: ## Migrate Up example: migrate-up database="postgres://username:password@host:port/db_name?sslmode=disable"
	@[ "${database}" ] || ( echo "database not set"; exit 1 )
	migrate -path database/migrations -database "$(database)" -verbose up $(step)

migrate-down: ## Migrate Down example: migrate-down database="postgres://username:password@host:port/db_name?sslmode=disable"
	@[ "${database}" ] || ( echo "database not set"; exit 1 )
	migrate -path database/migrations -database "$(database)" -verbose down $(step)