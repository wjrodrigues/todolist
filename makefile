pwd := $(shell pwd)

.SILENT: infra api start cover
.PHONY: infra api start cover

infra:
	docker network create todolist || true
	cp ${pwd}/docker/.env-example ${pwd}/docker/.env
	cd ${pwd}/docker && docker-compose up -d --build
	echo "Finish ✅"

api:
	cp ${pwd}/.env-example .env
	echo "Finish ✅"

cover:
	docker exec -u dev todolist-api sh -c "go test ./... -cover -coverprofile cover.out"
	docker exec todolist-api sh docker/ci/cover.sh

start:
	make infra
	make api
