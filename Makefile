.PHONY: gen-proto di serve clean-proto test export-local-env

ROOT := $(realpath .)


gen-proto: clean-proto
	protoc --proto_path=proto --twirp_out=:proto --go_out=:proto proto/*.proto

di:
	go generate di/wire_gen.go

dev-serve: export-local-env
	go run main.go

clean-proto:
	rm -rf proto/*.{pb,twirp}.go

test:
	go test ./...
	goapp test ./...

export-local-env:
	$(shell ln -fsn .env_local .env)
	cd _scripts && go build -o export_env export_env.go && cd ..
	$(ROOT)/_scripts/export_env
	rm $(ROOT)/_scripts/export_env
