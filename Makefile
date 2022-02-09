.PHONY: build

build:
	@echo "building .."
	@go build ${GCLFAGS} -o build/oasdiff main.go
	@ls -l build

build-debug:
	@echo "building for debug .."
	@go build -gcflags=all="-N -l" -o build/oasdiff main.go
	@ls -l build

clean:
	@rm -rf build
	@echo "cleanup done"

test:
	@echo "running tests .." 
	@go test ./... -timeout 30s -v -race
