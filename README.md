# OAS-DIFF
OAS 3.1 Validation and Diff Tool

## Dependencies
- Git
- Make
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
## Options
- Validate
    ````
   --base-file value  Path of the base OAS 3.1 file
   --help, -h         show help (default: false)
   ````
- Diff
    ````
    --base-file value       Path of the base OAS 3.1 file
    --second-file value     Path of the second OAS 3.1 file
    --html                  save the changelog file as a html report (default: false)
    --loose                 loosely diff (default: false)
    --include-file-path     Whether or not to include the full file path from the diff changelog (default: false)
    --exclude-descriptions  Whether or not to exclude descriptions from the diff changelog (default: false)
    --help, -h              show help (default: false)
    ````
## Examples
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
## Changelog Rules
- Arrays
    ````
    CREATE/DELETE -> Always the entire element
    UPDATE -> Only the property, exception if the property is the identifier
    ````
