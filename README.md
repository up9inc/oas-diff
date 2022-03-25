[![acceptance tests](https://github.com/up9inc/oas-diff/actions/workflows/acceptance_tests.yml/badge.svg?branch=develop)](https://github.com/up9inc/oas-diff/actions/workflows/acceptance_tests.yml)
# OAS-DIFF 
OAS 3.1 Validation and Diff Tool

## Dependencies
- Git
- Make
- Go 1.18+

## Setup
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
    ./build/oasdiff --version
    ````
## Options
- Validate
    ````
   --base-file value  Path of the base OAS 3.1 file
   --help, -h         show help (default: false)
   ````
- Diff
    ````
   --base-file value       path of the base OAS 3.1 file
   --second-file value     path of the second OAS 3.1 file
   --html                  save an html report (default: false)
   --loose                 loosely diff, ignores global case sensitivity for strings comparisons and ignore headers that start with 'x-' and 'user-agent' (default: false)
   --include-file-path     whether or not to include the full file path from the diff changelog (default: false)
   --exclude-descriptions  whether or not to exclude descriptions from the diff changelog (default: false)
   --help, -h              show help (default: false)
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
    ````
- Diff
    ````
    ./build/oasdiff diff --base-file examples/simple.json --second-file examples/simple2.json
    ./build/oasdiff diff --base-file examples/simple.json --second-file examples/simple2.json --loose
    ./build/oasdiff diff --base-file examples/simple.json --second-file examples/simple2.json --exclude-descriptions
    ./build/oasdiff diff --base-file examples/simple.json --second-file examples/simple2.json --loose --exclude-descriptions
    ./build/oasdiff diff --base-file examples/simple.json --second-file examples/simple2.json --html
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
- `Specification Extensions` are not suppoerted (TODO)
- `webhooks` field is only supoporting the `Path Item Object`, but it could also be a `Reference Object` (TODO)
- `openapi` field is not a part of the changelog, we only support OAS version 3.1, so any changes will cause a validation failure
- `jsonSchemaDialect` field is not a part of the changelog, we only support OAS version 3.1 and it uses `JSON Schema Validation Draft 2020-12`
