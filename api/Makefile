.PHONY: dev-up dev-down dev start-api test docs help

dev-up: 
	podman-compose up -d

dev-down:
	podman-compose down

start:
	go run ./cmd/main.go

docs:
	swag init --generalInfo ./cmd/main.go   

dev: docs start

test: 
	go test ./...

help:
	@echo "make dev-up 開啟開發場域"
	@echo "make dev-down 關閉開發場域"
	@echo "make start 啟動API"
	@echo "make dev 啟動 Dev API"
	@echo "make docs 產生Swagger API文件"