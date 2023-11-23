.PHONY: migrate-up

migrate-up:
	migrate -path internal/database/migration -database "postgresql://postgres:Docker@localhost:5432/postgres?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/database/migration -database "postgresql://postgres:Docker@localhost:5432/postgres?sslmode=disable" -verbose down