display-images:
	docker images | grep tmp-app

build-frontend:
	docker build -t tmp-app-frontend:latest ./frontend

run-frontend:
	docker run --rm -p 5173:5173 tmp-app-frontend
