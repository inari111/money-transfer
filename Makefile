.PHONY: gen-proto di dev-serve prod-serve clean-proto test export-local-env export-prod-env

ROOT := $(realpath .)


gen-proto: clean-proto
	protoc --proto_path=proto --twirp_out=:proto --go_out=:proto proto/*.proto

di:
	go generate di/wire_gen.go

dev-serve: export-local-env
	go run main.go

prod-serve: export-prod-env
	go run main.go

clean-proto:
	rm -rf proto/*.{pb,twirp}.go

test:
	go test ./...
	goapp test ./...

export-local-env:
	$(shell ln -fsn .env_local .env)
	sh $(ROOT)/_scripts/export_env.sh

export-prod-env:
	$(shell ln -fsn .env_prod .env)
	sh $(ROOT)/_scripts/export_env.sh
