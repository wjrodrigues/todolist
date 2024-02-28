pwd := $(shell pwd)

.SILENT: infra api start
.PHONY: infra api start

infra:
	docker network create todolist || true
	cp ${pwd}/docker/.env-example ${pwd}/docker/.env
	cd ${pwd}/docker && docker-compose up -d --build
	echo "Finish ✅"

api:
	cp ${pwd}/.env-example .env
	echo "Finish ✅"

start:
	make infra
	make api
