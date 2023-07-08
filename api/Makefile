.PHONY: dev dev-down start-api docs help

dev: 
	podman-compose up -d

dev-down:
	podman-compose down

docs:
	swag init --generalInfo ./cmd/main.go   

help:
	@echo "make dev 開啟開發場域"
	@echo "make dev-down 關閉開發場域"
	@echo "make docs 產生Swagger API文件"