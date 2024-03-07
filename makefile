pwd := $(shell pwd)

.SILENT: infra api start cover migrate-up migrate-down
.PHONY: infra api start cover migrate-up migrate-down

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
	docker exec todolist-migrate sh -c 'migrate -path ./migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}?sslmode=disable" -verbose down -all'

	echo "Migrate down test 🚧"
	docker exec todolist-migrate sh -c 'migrate -path ./migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}_test?sslmode=disable" -verbose down -all'
	echo "Finish ✅"

start:
	make infra
	make api
	make migrate-up
