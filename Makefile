run:
	go run main.go


POSTGRES_USER=postgres
POSTGRES_PASSWORD=12345
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DATABASE=postgres


DB_URL=postrges://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable


hello:
	echo "hello ${name}"


migrate-create:
	@read -p "Enter migration name: " name; \
	$(MIGRATE_BIN) create -ext sql -dir $(MIGRATIONS_DIR) $$name