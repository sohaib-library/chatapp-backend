include .env

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATION_DIR=./database/migration

.PHONY: migrate-up migrate-down migrate-status create-migration run docker-up docker-down docker-logs

migrate-up:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" down

migrate-status:
	goose -dir $(MIGRATION_DIR) postgres "$(DB_URL)" status

create-migration:
	goose -dir $(MIGRATION_DIR) create $(name) sql

run:
	go run main.go

docker-up:
	docker compose up --build -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f api

# ── Kubernetes (k3d) ──────────────────────────────────────────────────────────

k8s-cluster:
	k3d cluster create chatapp --port "8000:30800@loadbalancer" --port "30083:30083@loadbalancer"

k8s-load:
	docker build -t chatapp-backend:latest .
	k3d image import chatapp-backend:latest -c chatapp

k8s-deploy:
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -R -f k8s/

k8s-redeploy:
	docker build -t chatapp-backend:latest .
	k3d image import chatapp-backend:latest -c chatapp
	kubectl rollout restart deployment/chatapp-api -n chatapp

k8s-status:
	kubectl get all -n chatapp

k8s-logs:
	kubectl logs -f deployment/chatapp-api -n chatapp

k8s-down:
	k3d cluster delete chatapp

# ── Kubernetes (Minikube) ─────────────────────────────────────────────────────

minikube-start:
	minikube start --driver=docker --cpus=4 --memory=4096 --ports=8000:30800,30083:30083

minikube-load:
	docker build -t chatapp-backend:latest .
	minikube image load chatapp-backend:latest

minikube-deploy:
	kubectl apply -f k8s/namespace.yaml
	kubectl apply -R -f k8s/

minikube-redeploy:
	docker build -t chatapp-backend:latest .
	minikube image load chatapp-backend:latest
	kubectl rollout restart deployment/chatapp-api -n chatapp

minikube-status:
	kubectl get all -n chatapp

minikube-logs:
	kubectl logs -f deployment/chatapp-api -n chatapp

minikube-tunnel:
	@echo "Starting Minikube tunnel (requires sudo, keep running in separate terminal)"
	minikube tunnel

minikube-stop:
	minikube stop

minikube-delete:
	minikube delete