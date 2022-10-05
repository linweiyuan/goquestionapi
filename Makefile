up:
	migrate -path sqlc/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=question" -verbose up

down:
	migrate -path sqlc/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=question" -verbose down

gen:
	sqlc generate

.PHONY: up down gen
