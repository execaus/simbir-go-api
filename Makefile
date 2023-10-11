sqlc:
	sqlc generate

docker_postgres_build:
	docker build \
    		-t it_planeta_postgres \
    		-f ./docker/Dockerfile_postgres .

docker_api_build:
	docker build \
		-t it_planeta_api \
		-f ./docker/Dockerfile_api .

restart_test:
	docker rm 2023-it-planeta-web-api-webapi-1
	docker rm 2023-it-planeta-web-api-database-1
	make docker_postgres_build
	make docker_api_build
	docker compose up

lint:
	golangci-lint run -c ./.golangci.yml

update_docs:
	$HOME/go/bin/swag init -g cmd/main.go