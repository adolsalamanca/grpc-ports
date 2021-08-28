.PHONY: test up build run stop down docker-login image-push

test:
	cd server && go test ./... -race -v
	cd client && go test ./... -race -v

up:
	docker-compose -f docker-compose.yml up -d

build:
	docker-compose -f docker-compose.yml up -d --force-recreate --remove-orphans --build

stop:
	docker-compose stop

down:
	docker-compose -f tests/docker-compose.yml kill
	docker-compose -f tests/docker-compose.yml rm -sf

docker-login:
	docker login -u=${DOCKER_USERNAME} -p=${DOCKER_PASSWORD} ${DOCKER_REGISTRY}

image-push: docker-login
	docker push ${IMAGE_NAME}:${VERSION}
	docker push ${IMAGE_NAME}:latest