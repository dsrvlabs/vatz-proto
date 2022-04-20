help:
	@echo "Make Options:"
	@echo " go                  - Generate all : ${TARGET}"
	@echo " clean               - Clean up dist directory"

go:
	protoc -I=. \
	    --go_out ./ --go_opt paths=import \
	    --go-grpc_out ./ --go-grpc_opt paths=import \
		proto/**/**/**/*.proto

clean:
	$(call banner, "Clean up dist directory")
	rm -rf dist
