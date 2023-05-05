run:
	go run ./cmd/main.go

docker:
	docker run --name=task_db -e POSTGRES_PASSWORD=1234 -e POSTGRES_DB=task_db -p 5432:5432 -d --rm postgres

migrate_up:
	migrate -path ./schema -database 'postgres://postgres:1234@localhost:5432/task_db?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:1234@localhost:5432/task_db?sslmode=disable' down