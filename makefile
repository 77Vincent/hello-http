build:
	docker-compose build hello-http

run: build
	docker-compose up hello-http

push:
	docker build -t hello-http:latest .
	docker tag hello-http:latest frwentianqi/hello-http:latest
	docker push frwentianqi/hello-http:latest
	docker image rm hello-http:latest
