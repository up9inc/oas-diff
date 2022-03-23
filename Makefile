.PHONY: build

build:
	@echo "building oas-diff.."
	@go build -race ${GCLFAGS} -o ./build oasdiff.go
	@echo "binary saved to ./build folder"

build-debug:
	@echo "building oas-diff for debug .."
	@go build -race -gcflags=all="-N -l" -o ./build oasdiff.go
	@echo "binary saved to ./build folder"

clean:
	@rm -rf build
	@echo "cleanup done"

test:
	@echo "running tests .." 
	@go test ./... -timeout 30s -v -race

acceptance-test:
	@echo "running acceptance tests .."
	@cd acceptanceTests && $(MAKE) test
