# make install-frontend PKG=@tailwind/vite DEV=1
install-frontend:
	npm --prefix frontend install $(if $(DEV),-D) $(PKG)

display-images:
	docker images | grep tmp-app

build-frontend:
	docker build -t tmp-app-frontend:latest ./frontend

run-frontend:
	docker run --rm \
		-p 5173:5173 \
		-v "$(CURDIR)/frontend:/app" \
		-v /app/node_modules \
		tmp-app-frontend

build-backend:
	docker build -t tmp-app-backend:latest ./backend

run-backend:
	docker run --rm \
		-p 3000:3000 \
		-v "$(CURDIR)/backend:/app" \
		tmp-app-backend
