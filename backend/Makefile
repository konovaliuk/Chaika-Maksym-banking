.PHONY: help
help:
	go run cmd/*.go

.PHONY: migrate
migrate:
	BUNDEBUG=2 go run cmd/*.go db migrate

.PHONY: create_migration
create_migration:
	go run cmd/*.go db create ${name}

.PHONY: lint
lint:
	golangci-lint run -c ./golangci.yaml -v

.PHONY: fix_lint
fix_lint:
	golangci-lint run -c ./golangci.yaml -v --fix

.PHONY: tests
tests:
	go test -v ./...

.PHONY: gen_stubs
gen_stubs:
	protoc -I api/ --go_out=backend/api/ api/**/*.proto
