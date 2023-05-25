up:
	docker-compose up --build -d
stop:
	docker-compose stop
migration:
ifndef NAME
	$(error NAME is undefined. Pass it like this: make migration NAME=my_migration)
endif
	@migrate create -ext sql -dir migrations/ -seq $(NAME)
migrate-up:
	@migrate -path migrations/ -database "postgresql://postgres:password@localhost:5432/price_checker_db?sslmode=disable" -verbose up
migrate-down:
	@migrate -path migrations/ -database "postgresql://postgres:password@localhost:5432/price_checker_db?sslmode=disable" -verbose down
migration-fix:
ifndef VERSION
	$(error VERSION is undefined. Pass it like this: make migration-fix VERSION=00001)
endif
	@migrate -path migrations/ -database "postgresql://postgres:password@localhost:5432/price_checker_db?sslmode=disable" force $(VERSION)
help:
	@echo "Available targets:"
	@echo "  up:                    Starts the containers defined in docker-compose.yml."
	@echo "  stop:                  Stops the containers."
	@echo "  migration:             Creates a new migration with the name passed in the NAME variable."
	@echo "                          Example: make migration NAME=my_migration"
	@echo "  migrate-up:            Runs all the pending migrations in the database."
	@echo "  migrate-down:          Reverts the last migration applied to the database."
	@echo "  migration-fix:         Forces the database version to match the one specified in the VERSION variable."
	@echo "                          Example: make migration-fix VERSION=00001"
