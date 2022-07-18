.PHONY: build

run:
	docker-compose up --build

.DEFAULT_GOAL := run