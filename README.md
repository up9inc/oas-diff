[![acceptance tests](https://github.com/up9inc/oas-diff/actions/workflows/acceptance_tests.yml/badge.svg?branch=develop)](https://github.com/up9inc/oas-diff/actions/workflows/acceptance_tests.yml)
# OAS-DIFF 
OAS 3.1 Validation and Diff Tool

## Dependencies
- Git
- Make
- Go 1.18+

## Build
- Build
    ````
    make build
    ````
- Run
    ````
    ./build/oasdiff --version
    ````
## Options
- Validate
    ````
    --base-file value, --f1 value    path of the base OAS 3.1 file
    --help, -h                                   show help (default: false)
   ````
- Diff
    ````
    --base-file value, --f1 value    path of the base OAS 3.1 file
    --second-file value, --f2 value  path of the second OAS 3.1 file
    --type-filter value, --tf value  changelog Type filter (create/update/delete)
    --output-html, --oh              save an html report (default: false)
    --output-endpoint, --oe          endpoint based changelog output (default: false)
    --loose, -l                      loosely diff, ignores global case sensitivity for strings comparisons and ignore headers that start with 'x-' and 'user-agent' (default: false)
    --include-file-path, --ifp       whether or not to include the full file path from the diff changelog (default: false)
    --ignore-descriptions, --id      whether or not to ignore descriptions when performing the diff (default: false)
    --ignore-examples, --ie          whether or not to ignore examples when performing the diff (default: false)
    --help, -h                       show help (default: false)
    ````
## Examples
- Version
    ````
    ./build/oasdiff -v
    ````
- Available Commands
    ````
    ./build/oasdiff
    ````
- Validate
    ````
    ./build/oasdiff validate --base-file examples/invalid.json
    ./build/oasdiff validate --f1 examples/invalid.json
    ````
- Diff
    ````
    ./build/oasdiff diff --base-file examples/simple.json --second-file examples/simple2.json
    ./build/oasdiff diff --base-file examples/simple.json --second-file examples/simple2.json --loose
    ./build/oasdiff diff --f1 examples/simple.json --f2 examples/simple2.json --type-filter create
    ./build/oasdiff diff --f1 examples/simple.json --f2 examples/simple2.json --type-filter update
    ./build/oasdiff diff --f1 examples/simple.json --f2 examples/simple2.json --type-filter delete
    ./build/oasdiff diff --f1 examples/simple.json --f2 examples/simple2.json --ignore-descriptions
    ./build/oasdiff diff --f1 examples/simple.json --f2 examples/simple2.json --ignore-examples
    ./build/oasdiff diff --f1 examples/simple.json --f2 examples/simple2.json --loose --ignore-descriptions --ignore-examples
    ./build/oasdiff diff --f1 examples/simple.json --f2 examples/simple2.json --output-html
    ./build/oasdiff diff --f1 examples/simple.json --f2 examples/simple2.json --output-endpoint
    ./build/oasdiff diff --f1 examples/simple.json --f2 examples/simple2.json --output-endpoint --output-html --type-filter update
    ````

## Array Identifiers
 Array identifiers are used only for arrays to compare arrays by a matching identifier and not based on order. If an identifiable element is found in both the from and to structures, they will be directly compared

- Servers
    ````
    URL         string             `json:"url,omitempty" diff:"url,identifier"`
    ````
- Tags
    ````
    Name         string        `json:"name,omitempty" diff:"name,identifier"`
    ````
- Parameters
    ````
    Name         string        `json:"name,omitempty" diff:"name,identifier"`
    ````

## Changelog Rules
- Arrays
    ````
    CREATE/DELETE -> Always the entire element
    UPDATE -> Only the property, exception if the property is the identifier
    ````

## Limitations
- `openapi` field is not a part of the changelog, we only support OAS version 3.1, so any changes will cause a validation failure
- `jsonSchemaDialect` field is not a part of the changelog, we only support OAS version 3.1 and it uses `JSON Schema Validation Draft 2020-12`
- `Specification Extensions` are not supported (TODO)
- `parameters` array of `Reference Object` are not supported (TODO)