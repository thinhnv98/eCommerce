local:
	go run main.go

build:
	docker build --rm --force-rm -t e-commerce-api-image .
	#docker image prune --filter label=stage=Builder --filter label=stage=Final

run: build
	docker compose up

down:
	docker compose down