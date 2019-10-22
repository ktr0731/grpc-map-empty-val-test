SHELL := /bin/bash

export GO111MODULE := on

.PHONY: bin
bin: 
	@go build -o bin/server ./server
	@go build -o bin/client ./client

.PHONY: run/server/nil
run/server/nil: bin
	./bin/server -nil

.PHONY: run/server/nonnil
run/server/nonnil: bin
	./bin/server

.PHONY: run/grpcurl
run/grpcurl:
	@grpcurl --version
	grpcurl -plaintext :50051 api.Service.GetMap

.PHONY: run/evans
run/evans:
	@evans --version
	echo {} | evans -r --call GetMap

.PHONY: run/client
run/client: bin
	./bin/client
