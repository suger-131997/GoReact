.PHONY: dev build run gen-types gen-js-builder

# ホットリロードを有効にして起動
dev:
	air

gen-js-builder:
	go run cmd/gen-js-builder/main.go

gen-js: gen-js-builder
	go run gen/gen.go

gen-types:
	go run cmd/gen-types/main.go

build: gen-js
	go build -o tmp/main ./entrypoint/prod

run: build
	./tmp/main

