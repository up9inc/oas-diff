module github.com/up9inc/oas-diff

go 1.17

require (
	github.com/santhosh-tekuri/jsonschema/v5 v5.0.0
	github.com/tidwall/gjson v1.14.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	golang.org/x/net v0.0.0-20190603091049-60506f45cf65 // indirect
	google.golang.org/appengine v1.6.6 // indirect
)

require (
	github.com/r3labs/diff/v2 v2.14.8
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/urfave/cli/v2 v2.3.0
)

replace github.com/r3labs/diff/v2 v2.14.8 => ./lib
