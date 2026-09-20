# make install-frontend PKG=@tailwind/vite DEV=1
install-frontend:
	npm --prefix frontend install $(if $(DEV),-D) $(PKG)

display-images:
	docker images | grep tmp-app

build-frontend:
	docker build -t tmp-app-frontend:latest ./frontend

run-frontend: build-frontend
	docker run -d --rm \
		--name tmp-app-frontend \
		-p 5173:5173 \
		-v "$(CURDIR)/frontend:/app" \
		-v /app/node_modules \
		tmp-app-frontend

remove-frontend:
	docker rm -f tmp-app-frontend

build-backend:
	docker build -t tmp-app-backend:latest ./backend

run-backend: build-backend
	docker run -d --rm \
		--name tmp-app-backend \
		-p 3000:3000 \
		-v "$(CURDIR)/backend:/app" \
		tmp-app-backend

remove-backend:
	docker rm -f tmp-app-backend

kind-create:
	kind create cluster --config kind-config.yaml

kind-load: build-frontend build-backend
	kind load docker-image tmp-app-frontend:latest --name tmp-app
	kind load docker-image tmp-app-backend:latest --name tmp-app

kind-deploy: kind-load
	kubectl apply -k k8s
	kubectl rollout restart deployment/frontend deployment/backend -n tmp-app

kind-status:
	kubectl get pods,services -n tmp-app

kind-delete:
	kind delete cluster --name tmp-app
