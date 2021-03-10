build:
	docker-compose build
build-nocache:
	docker-compose build --no-cache
sync-dependencies:
	docker-compose up -d