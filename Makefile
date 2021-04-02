build:
	docker-compose build api-example

run:
	docker-compose up api-example

stop:
	docker-compose stop

test:
	go test  ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

swag: #todo doesn't work
	swag init -g cmd/main.go

