package model

// https://spec.openapis.org/oas/v3.1.0#security-scheme-object
type SecurityScheme struct {
	Type             string     `json:"type,omitempty" diff:"type"`
	Description      string     `json:"description,omitempty" diff:"description"`
	Name             string     `json:"name,omitempty" diff:"name"`
	In               string     `json:"in,omitempty" diff:"in"`
	Scheme           string     `json:"scheme,omitempty" diff:"scheme"`
	BearerFormat     string     `json:"bearerFormat,omitempty" diff:"bearerFormat"`
	Flows            OAuthFlows `json:"flows,omitempty" diff:"flows"`
	OpenIdConnectUrl string     `json:"openIdConnectUrl,omitempty" diff:"openIdConnectUrl"`

	// Reference object
	Ref     string `json:"$ref,omitempty" diff:"$ref"`
	Summary string `json:"summary,omitempty" diff:"summary"`
}

// https://spec.openapis.org/oas/v3.1.0#oauth-flow-object
type OAuthFlow struct {
	AuthorizationUrl string     `json:"authorizationUrl,omitempty" diff:"name,authorizationUrl"`
	TokenUrl         string     `json:"tokenUrl,omitempty" diff:"name,tokenUrl"`
	RefreshUrl       string     `json:"refreshUrl,omitempty" diff:"refreshUrl"`
	Scopes           StringsMap `json:"scopes,omitempty" diff:"scopes"`
}

// https://spec.openapis.org/oas/v3.1.0#oauth-flows-object
type OAuthFlows struct {
	Implicit          OAuthFlow `json:"implicit,omitempty" diff:"implicit"`
	Password          OAuthFlow `json:"password,omitempty" diff:"password"`
	ClientCredentials OAuthFlow `json:"clientCredentials,omitempty" diff:"clientCredentials"`
	AuthorizationCode OAuthFlow `json:"authorizationCode,omitempty" diff:"authorizationCode"`
}
