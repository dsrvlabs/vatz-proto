help:
	@echo "Make Options:"
	@echo " all                 - Generate all : compile all proto-buffers into Pbs"
	@echo " init                - Install required modules : install protoc-gen-go & protoc-gen-go-grpc"
	@echo " build               - Generate all : compile all protobuffers Pbs"
	@echo " clean               - Clean up compile directory"

all:
	make init
	make build

init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

build:
	protoc -I=. \
	    --go_out ./ --go_opt paths=import \
	    --go-grpc_out ./ --go-grpc_opt paths=import \
		proto/**/**/**/*.proto

compile:
	sudo protoc -I=. \
          --go_out ./ --go_opt paths=import \
          --go-grpc_out ./ --go-grpc_opt paths=import \
        proto/**/**/**/*.proto
clean:
	$(call banner, "Clean up compile directory")
	rm -rf manager plugin
