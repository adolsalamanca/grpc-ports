.PHONY: test up build run stop down

test: | up
	docker exec ports_client go test ./... -v
	docker exec ports_port go test ./... -v

up:
	docker-compose -f docker-compose.yml up -d

build:
	docker-compose -f docker-compose.yml up -d --force-recreate --build

stop:
	docker-compose stop

down:
	docker-compose down
