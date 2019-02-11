.PHONY: gen-proto di serve clean-proto test

gen-proto: clean-proto
	protoc --proto_path=proto --twirp_out=:proto --go_out=:proto proto/*.proto

di:
	go generate di/wire_gen.go

serve:
	go run main.go

clean-proto:
	rm -rf proto/*.{pb,twirp}.go

test:
	go test ./...
	goapp test ./...