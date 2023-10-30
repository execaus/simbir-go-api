sqlc:
	sqlc generate

docker_postgres_build:
	docker build \
    		-t simbir-go-postgres \
    		-f ./docker/Dockerfile_postgres .

docker_api_build:
	docker build \
		-t simbir-go-api \
		-f ./docker/Dockerfile_api .

lint:
	golangci-lint run -c ./.golangci.yml

update_docs:
	$HOME/go/bin/swag init -g cmd/main.go