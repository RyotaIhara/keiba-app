# --------------------------------------------------
# npm関連
# --------------------------------------------------
# make install-frontend PKG=@tailwind/vite DEV=1
install-frontend:
	npm --prefix frontend install $(if $(DEV),-D) $(PKG)

# --------------------------------------------------
# docker関連
# --------------------------------------------------
COMPOSE=docker compose --env-file .env.local

display-images:
	docker images | grep keiba-app

build-frontend:
	docker build -t keiba-app-frontend:latest ./frontend

run-frontend: build-frontend
	docker run -d --rm \
		--name keiba-app-frontend \
		-p 5173:5173 \
		-v "$(CURDIR)/frontend:/app" \
		-v /app/node_modules \
		keiba-app-frontend

remove-frontend:
	docker rm -f keiba-app-frontend

build-backend:
	docker build -t keiba-app-backend:latest ./backend

run-backend: build-backend
	docker run -d --rm \
		--name keiba-app-backend \
		-p 3000:3000 \
		-v "$(CURDIR)/backend:/app" \
		keiba-app-backend

remove-backend:
	docker rm -f keiba-app-backend

kind-load: build-frontend build-backend
	kind load docker-image keiba-app-frontend:latest --name keiba-app
	kind load docker-image keiba-app-backend:latest --name keiba-app

kind-deploy: kind-load
	kubectl apply -k k8s
	kubectl rollout restart deployment/frontend deployment/backend -n keiba-app

kind-status:
	kubectl get pods,services -n keiba-app

# --------------------------------------------------
# Kubernetis関連
# --------------------------------------------------
APP_NAME=keiba-app

kind-create:
	kind create cluster --name $(APP_NAME)

kind-delete:
	kind delete cluster --name $(APP_NAME)

pods:
	kubectl get pods

services:
	kubectl get svc

# backend（backend-upコマンドで問題なし）
BACKEND_IMAGE=keiba-app-backend:latest

backend-build:
	docker build -t $(BACKEND_IMAGE) ./backend
backend-load:
	kind load docker-image $(BACKEND_IMAGE) --name $(APP_NAME)
backend-deploy:
	kubectl apply -f k8s/backend-deployment.yaml
backend-service:
	kubectl apply -f k8s/backend-service.yaml
backend-restart:
	kubectl rollout restart deployment backend

## 起動
backend-up: backend-build backend-load backend-deploy backend-service backend-restart

## Serviceにアクセスできるようにする
backend-forward:
	kubectl port-forward svc/backend 3000:3000
## BackGround起動もできる
backend-forward-bg:
	kubectl port-forward svc/backend 5173:5173 > /tmp/backend-port-forward.log 2>&1 & echo $$! > /tmp/backend-port-forward.pid
backend-forward-bg-stop:
	kill `cat /tmp/backend-port-forward.pid`
	rm -f /tmp/backend-port-forward.pid

# frontend（frontend-upコマンドで問題なし）
FRONTEND_IMAGE=keiba-app-frontend:latest

frontend-build:
	docker build -t $(FRONTEND_IMAGE) ./frontend
frontend-load:
	kind load docker-image $(FRONTEND_IMAGE) --name $(APP_NAME)
frontend-deploy:
	kubectl apply -f k8s/frontend-deployment.yaml
frontend-service:
	kubectl apply -f k8s/frontend-service.yaml
frontend-restart:
	kubectl rollout restart deployment frontend

## 起動
frontend-up: frontend-build frontend-load frontend-deploy frontend-service frontend-restart

## Serviceにアクセスできるようにする
frontend-forward:
	kubectl port-forward svc/frontend 5173:5173
## BackGround起動もできる
frontend-forward-bg:
	kubectl port-forward svc/frontend 5173:5173 > /tmp/frontend-port-forward.log 2>&1 & echo $$! > /tmp/frontend-port-forward.pid
frontend-forward-bg-stop:
	kill `cat /tmp/frontend-port-forward.pid`
	rm -f /tmp/frontend-port-forward.pid

# 削除
deployment-delete:
	kubectl delete deployment frontend
	kubectl delete deployment backend
service-delete:
	kubectl delete service frontend
	kubectl delete service backend

# --------------------------------------------------
# DB関連
# --------------------------------------------------
.PHONY: sql

## SQLファイルをMySQLコンテナへ投入
## 例: make sql SQL=sql/schema/create_user.sql
sql:
	@test -n "$(SQL)" || (echo "使用例: make sql SQL=path/to/file.sql" >&2; exit 1)
	@test -f "$(SQL)" || (echo "SQLファイルが見つかりません: $(SQL)" >&2; exit 1)
	$(COMPOSE) exec -T mysql \
		sh -c 'mysql --default-character-set=utf8mb4 -u"$$MYSQL_USER" -p"$$MYSQL_PASSWORD" "$$MYSQL_DATABASE"' \
		< "$(SQL)"

sql-init:
	$(MAKE) sql SQL=sql/schema/create_user.sql
	$(MAKE) sql SQL=sql/schema/create_race_course.sql
	$(MAKE) sql SQL=sql/schema/create_race.sql
	$(MAKE) sql SQL=sql/data/insert_user_data.sql
	$(MAKE) sql SQL=sql/data/insert_race_course_data.sql
	$(MAKE) sql SQL=sql/data/insert_race_data.sql
