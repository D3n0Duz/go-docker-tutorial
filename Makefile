run:
	go run *.go

build:
	go build -o main

docker-build:
	GOOS=linux GOARCH=amd64 go build -o main
	docker build -t helloworld-go .

docker-run:
	docker run -it -p 8080:8080 helloworld-go

docker-push:
	docker push helloworld-go
