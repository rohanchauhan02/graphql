.PHONY: run gen

run:
	go run cmd/app/main.go

gen:
	gqlgen generate



