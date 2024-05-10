API_PATH=api/weather
PROTO_OUT_DIR=pkg/weatherapi
PROTO_API_DIR=$(API_PATH)
ARGS=-fix

.PHONY: gen
gen: gen-proto generate

.PHONY: gen-proto
gen-proto:
	protoc \
    	-I $(API_PATH)/v1 \
    	--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
    	--go-grpc_out=$(PROTO_OUT_DIR)  --go-grpc_opt=paths=source_relative \
    ./$(PROTO_API_DIR)/v1/*.proto

evans:
	evans --port 9001 -r repl

test:
	go test ./... -cover -count=1

generate:
	go generate ./...

lint:
	go/lint proto/lint

go/lint:
	golangci-lint run  --config=.golangci.yml --timeout=30s ./...

proto/lint:
	protolint lint $(ARGS) $(PROTO_API_DIR)/*

run:
	go run cmd/weather/main.go

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root weather

dropdb:
	docker exec -it postgres12 dropdb ozinshe

migrateup:
	migrate -database 'postgresql://root:123@localhost:5432/weather?sslmode=disable' -path scheme up

migratedown:
	migrate -database 'postgresql://root:123@localhost:5432/weather?sslmode=disable' -path scheme down

fixdirty:
	migrate -database 'postgresql://root:123@localhost:5432/weather?sslmode=disable' -path migrations force 1 

psql:
	docker exec -it postgres12  psql -U root -d weather