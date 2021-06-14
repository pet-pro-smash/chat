.PHONY: build run migrate_up
build:
	go build -o build/bin/main cmd/main.go

run:
	./build/bin/main

migrate_up:
	migrate -path ./schema/ -database "postgres://postgres:${POSTGRES_PASSWORD}@localhost:5432/chat-db?sslmode=disable" up

migrate_down:
	migrate -path ./schema/ -database "postgres://postgres:${POSTGRES_PASSWORD}@localhost:5432/chat-db?sslmode=disable" down
