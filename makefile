pwd := $(shell pwd)

.SILENT: infra api start cover migrate-up migrate-down cover-html
.PHONY: infra api start cover migrate-up migrate-down cover-html

infra:
	docker network create todolist || true
	cp ${pwd}/docker/.env-example ${pwd}/docker/.env
	cd ${pwd}/docker && docker-compose up -d --build todolist-api todolist-db
	echo "Finish ✅"

api:
	cp ${pwd}/.env-example .env
	echo "Finish ✅"

cover:
	docker exec -u dev todolist-api sh -c "go test ./... -cover -coverprofile cover.out"
	docker exec todolist-api sh docker/ci/cover.sh

cover-html:
	docker exec -u dev todolist-api sh -c "go test ./... -cover -coverprofile cover.out"
	docker exec todolist-api sh -c "go tool cover -html=cover.out -o coverage/index.html"

migrate-up:
	echo "Migrate up dev 🛠️"
	docker compose -f docker/docker-compose.yml up -d todolist-migrate
	docker exec todolist-migrate sh -c 'migrate -path ./migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}?sslmode=disable" -verbose up'

	echo "Migrate up test 🚧"
	docker exec todolist-migrate sh -c 'migrate -path ./migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}_test?sslmode=disable" -verbose up'
	echo "Finish ✅"

migrate-down:
	echo "Migrate down dev 🛠️"
	docker compose -f docker/docker-compose.yml up -d todolist-migrate
	docker exec todolist-migrate sh -c 'migrate -path ./migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}?sslmode=disable" -verbose down 1'

	echo "Migrate down test 🚧"
	docker exec todolist-migrate sh -c 'migrate -path ./migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}_test?sslmode=disable" -verbose down 1'
	echo "Finish ✅"

start:
	make infra
	make api
	make migrate-up
