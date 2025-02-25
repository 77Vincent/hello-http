build:
	docker-compose build hello-http

run: build
	docker-compose up hello-http

push:
	docker build -t hello-http:latest .
	docker tag hello-http:latest frwentianqi/hello-http:0.1.0
	docker push frwentianqi/hello-http:0.1.0
	docker image rm hello-http:latest
