build:
	docker-compose build

run:
	docker-compose up

migrate-up:
	migrate -path ./schema -database 'postgres://dev:dev@localhost:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://dev:dev@localhost:5432/postgres?sslmode=disable' down