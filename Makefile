gen:
	sqlboiler --wipe psql

down: down-db down-api

down-db:
	docker compose down e-commerce-postgres-db

down-api:
	docker compose down e-commerce-api

build-db:
	docker compose build e-commerce-postgres-db

setup: down-db build-db
	docker compose up e-commerce-postgres-db

build-api:
	docker build --rm --force-rm -t e-commerce-api-image .
	#docker image prune --filter label=stage=Builder --filter label=stage=Final

api:
	docker compose up e-commerce-api

run: down build-db build-api
	docker compose up