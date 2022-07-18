.PHONY: build

run:
	go run ./cmd/gophermart/main.go

.DEFAULT_GOAL := run