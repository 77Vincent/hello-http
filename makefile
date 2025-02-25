build:
	docker-compose build hello-http

run: build
	docker-compose up hello-http

push:
	docker buildx build --platform linux/amd64,linux/arm64 -t frwentianqi/hello-http:latest -t frwentianqi/hello-http:0.1.0 --push .
