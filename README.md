# oas-diff
OAS 3.1 Validation and Diff Tool

## Requisits
- Go 1.17+

## Run
- Build
    ````
    make build
    ````
- Run
    ````
    ./build/oasdiff --help
    ````

## Examples
- Validate
    ````
    ./build/oasdiff validate --file test/shipping_invalid.json
    ````
- Diff
    ````
    ./build/oasdiff diff --file test/shipping_valid.json --file2 test/shipping_invalid.json
    ````
