run:
	go run main.go


POSTGRES_USER=postgres
POSTGRES_PASSWORD=12345
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DATABASE=postgres


DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable
#url=postgres://postgres:12345@localhost:5432/postgres?sslmode=disable

hello:
	echo "hello ${name}"


migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir ./migrations  $(name)



version:
	migrate -database "$(DB_URL)" -path ./migrations version
migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

migrate-reset:
	migrate -path ./migrations -database "$(DB_URL)" down -all



