swagger:
	swag init -g cmd/main.go

build:
	go build -o todo-list ./cmd

run: build
	./todo-list

test:
	go test ./tests	

docker-build:
	docker-compose build	

docker-run: docker-build
	docker-compose up