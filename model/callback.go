package model

// https://spec.openapis.org/oas/v3.1.0#callback-object
type Callbacks map[string]*Callback
type Callback map[string]*PathItem