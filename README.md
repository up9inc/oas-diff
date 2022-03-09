# oas-diff
OAS 3.1 Validation and Diff Tool

## Requisits
- Go 1.17+

## Run
- Setup (Just once)
    ````
    ./setup.sh
    ````
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
    ./build/oasdiff diff --file test/simple.json --file2 test/simple2.json
    ````
## Changelog Rules
- Arrays
    ````
    CREATE/DELETE -> Always the entire element
    UPDATE -> Only the property, exception if the property is the identifier
    ````
