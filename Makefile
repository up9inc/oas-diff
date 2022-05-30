.PHONY: build

generate:
	@echo "running go generate .."
	@go generate ./embed
	@echo "files added to embed box"

template:
	@echo "building template file .."
	@(cd ui; npm i ; npm run build; )
	@echo "template file saved to ./static folder"

build:
	$(MAKE) template
	$(MAKE) generate
	@echo "building oas-diff .."
	@mkdir -p build
	@go build -race ${GCLFAGS} -o ./build oasdiff.go
	@echo "binary saved to ./build folder"

build-debug:
	@echo "building oas-diff for debug .."
	@mkdir build -p
	@go build -race -gcflags=all="-N -l" -o ./build oasdiff.go
	@echo "binary saved to ./build folder"

clean:
	@rm -rf build
	@rm -rf static
	@rm -rf embed/blob.go
	@echo "cleanup done"

test:
	@echo "running tests .." 
	@go test ./... -timeout 30s -v -race

acceptance-test:
	@echo "running acceptance tests .."
	@cd acceptanceTests && $(MAKE) test
