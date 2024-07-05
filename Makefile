dep:
	go mod tidy

run-user:
	go run cmd/time-tracker/main.go

test:
	go test -short -cover ./...

build-user:
	go build -o bin/server cmd/time-tracker/main.go

docker-image:
	docker build -t server:v1 .

docker-run:
	docker run -it -d -p 3000:3000 --name server
