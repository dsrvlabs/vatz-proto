help:
	@echo "Make Options:"
	@echo " go                  - Generate all : ${TARGET}"
	@echo " clean               - Clean up dist directory"

go:
	rm -rf dist
	mkdir dist
	protoc -I=. \
	    --go_out dist/ --go_opt paths=source_relative \
	    --go-grpc_out dist/ --go-grpc_opt paths=source_relative \
		proto/**/**/**/*.proto

clean:
	$(call banner, "Clean up dist directory")
	rm -rf dist
