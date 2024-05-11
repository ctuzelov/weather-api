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
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=postgres --owner=postgres postgres

dropdb:
	docker exec -it postgres12 dropdb weather

migrateup:
	migrate -database 'postgresql://postgres:123@localhost:5432/postgres?sslmode=disable' -path scheme/migrations up

migratedown:
	migrate -database 'postgresql://postgres:123@localhost:5432/postgres?sslmode=disable' -path scheme/migrations down

fixdirty:
	migrate -database 'postgresql://postgres:123@localhost:5432/postgres?sslmode=disable' -path scheme/migrations force 1 

psql:
	docker exec -it postgres12  psql -U root -d postgres