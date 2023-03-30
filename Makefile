hello:
	echo "Hello"

build:
	go build main.go

run:
	go run main.go

mysql:
	docker compose up
