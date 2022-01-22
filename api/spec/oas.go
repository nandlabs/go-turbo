package spec

//OAS  as specified by OAS version 3.1.0 https://spec.openapis.org/oas/v3.1.0
//This is the top level item in the project items
type OAS struct {
	//Openapi version
	OpenAPI string `json:"openapi" yaml:"openapi"` // Required
	Info    Info   `json:"info" yaml:"info"`       // Required

	//The default value for the $schema keyword within Schema Objects contained within this OAS document.
	//This MUST be in the form of a URI.
	JsonSchemaDialect string                 `json:"jsonSchemaDialect,omitempty" yaml:"jsonSchemaDialect,omitempty"`
	Servers           []*Server              `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths             map[string]*PathItem   `json:"paths" yaml:"paths"` // Required
	Webhooks          map[string]*PathItem   `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
	Components        *Components            `json:"components,omitempty" yaml:"components,omitempty"`
	Security          []*SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
	Tags              []*Tag                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs      *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

//Info as specified by OAS version 3.1.0 https://spec.openapis.org/oas/v3.1.0#info-object
type Info struct {
	Title          string  `json:"title" yaml:"title"` // Required
	Summary        string  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description    string  `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        string  `json:"version" yaml:"version"` // Required
}

//Contact as specified by OAS version 3.1.0 https://spec.openapis.org/oas/v3.1.0#contact-object
type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

//License as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#license-object
type License struct {
	Name       string `json:"name" yaml:"name"`                                 // Required
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"` // Required
	URL        string `json:"url,omitempty" yaml:"url,omitempty"`
}

//Tag as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#tag-object
type Tag struct {
	Name         string                `json:"name" yaml:"name"` // Required
	Description  string                `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// Server as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#server-object
type Server struct {
	URL         string                    `json:"url" yaml:"url"` //Required
	Description string                    `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

// ServerVariable as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#server-variable-object
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitempty"` //Required
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

//PathItem as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#path-item-object
type PathItem struct {
	Reference
	Summary     string      `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Get         *Operation  `json:"get,omitempty" yaml:"get,omitempty"`
	Put         *Operation  `json:"put,omitempty" yaml:"put,omitempty"`
	Post        *Operation  `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      *Operation  `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options     *Operation  `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *Operation  `json:"head,omitempty" yaml:"head,omitempty"`
	Patch       *Operation  `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace       *Operation  `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []Server    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

//Operation as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#operation-object
type Operation struct {
	Tags         []string               `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  string                 `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []*Parameter           `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBody           `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    map[string]*Response   `json:"responses,omitempty" yaml:"responses,omitempty"`
	Callbacks    map[string]*Callback   `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []*SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []*Server              `json:"servers,omitempty" yaml:"servers,omitempty"`
}

//SecurityScheme as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#security-scheme-object
type SecurityScheme struct {
	Type             string     `json:"type" yaml:"type"` //Required Enum Values ( apiKey,http,oauth2,openIdConnect)
	Description      string     `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string     `json:"name" yaml:"name"`     //Required
	In               string     `json:"in" yaml:"in"`         //Required Valid values are "query", "header" or "cookie".
	Scheme           string     `json:"scheme" yaml:"scheme"` //Required
	BearerFormat     string     `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            OauthFlows `json:"flows" yaml:"flows"`                       //Required
	OpenIDConnectURL string     `json:"openIdConnectUrl" yaml:"openIdConnectUrl"` //Required
}

//SecurityRequirement as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#security-requirement-object
type SecurityRequirement struct {
	Fields map[string][]string
}

// OauthFlows as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#oauth-flow-object
type OauthFlows struct {
	Implicit          OauthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          OauthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials OauthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode OauthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

// OauthFlow as specified by OAS version 3.0.3 https://spec.openapis.org/oas/v3.1.0#oauth-flow-object
type OauthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl" yaml:"authorizationUrl"` //Required
	TokenURL         string            `json:"tokenUrl" yaml:"tokenUrl"`                 //Required
	RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"` //Required

}

//Parameter as specified by OAS Version 3.0.3 https://spec.openapis.org/oas/v3.1.0#parameter-object
type Parameter struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	//Valid values are "query", "header", "path" or "cookie"
	In              string `json:"in" yaml:"in"` //Required
	Description     string `json:"description,omitempty" yaml:"description,omitempty" `
	Required        bool   `json:"required" yaml:"required"`
	Deprecated      bool   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

	// Valid Values for style are
	// 1. matrix
	// 2. label
	// 3. form
	// 4. simple
	// 5. spaceDelimited
	// 6. pipeDelimited
	// 7. deepObject
	// Default  values  based on the value of in field
	// in=query or in=cookie => style=form,
	// in=path or in=header  => style=simple
	Style         string `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool   `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool   `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`

	//Example and Examples are mutually exclusive
	Example
	Examples map[string]Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	//Schema and content Type are mutually exclusive
	Content map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Schema  Schema               `json:"schema,omitempty" yaml:"schema,omitempty"`
}

//Header type as specified by OAS version 3.0.3
type Header struct {
	//	Either Ref or Name will be present.
	Reference
	Description     string `json:"description,omitempty" yaml:"description,omitempty" `
	Required        bool   `json:"required" yaml:"required"`
	Deprecated      bool   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

	// Valid Values for style are
	// 1. matrix
	// 2. label
	// 3. form
	// 4. simple
	// 5. spaceDelimited
	// 6. pipeDelimited
	// 7. deepObject
	// Default  values  based on the value of in field
	// in=query or in=cookie => style=form,
	// in=path or in=header  => style=simple
	Style         string `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool   `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool   `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`

	//Example and Examples are mutually exclusive
	Example
	Examples map[string]Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	//Schema and content Type are mutually exclusive
	Content map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Schema  Schema               `json:"schema,omitempty" yaml:"schema,omitempty"`
}

//Schema Object

//RequestBody object as per https://spec.openapis.org/oas/v3.1.0#requestBodyObject
type RequestBody struct {
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Required    bool                 `json:"required,omitempty" yaml:"required,omitempty"`
}

//Response Object as per https://spec.openapis.org/oas/v3.1.0#requestBodyObject
type Response struct {
	Reference
	Description string               `json:"description" yaml:"description"`
	Headers     map[string]Header    `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]Link      `json:"links,omitempty" yaml:"links,omitempty"`
}

//The Link object as per OAS 3.0.3 https://spec.openapis.org/oas/v3.1.0#link-object
type Link struct {
	Reference
	OperationRef string                 `json:"operationRef,omitempty," yaml:"operationRef,omitempty"`
	OperationID  string                 `json:"operationId,omitempty," yaml:"operationId,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty," yaml:"parameters,omitempty"`
	RequestBody  interface{}            `json:"requestBody,omitempty," yaml:"requestBody,omitempty"`
	Description  string                 `json:"description,omitempty," yaml:"description,omitempty"`
	Server       Server                 `json:"server,omitempty," yaml:"server,omitempty"`
}

//Callback struct to hold the
type Callback struct {
	Reference
	Callbacks map[string]PathItem
}

//Components object as specified by OAS version 3.0.3  https://spec.openapis.org/oas/v3.1.0#components-object
type Components struct {
	Schemas         map[string]*Schema         `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]*Response       `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]*Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]*Example        `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]*RequestBody    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]*Header         `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           map[string]*Link           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]*Callback       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	PathItems       map[string]*PathItem       `json:"pathItems,omitempty" yaml:"pathItems,omitempty"`
}

//MediaType object  as per OAS 3.0.3
type MediaType struct {
	Schema Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example
	Examples map[string]Example  `json:"examples,omitempty" yaml:"examples,omitempty"` //Example and Examples are mutually exclusive
	Encoding map[string]Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

//Encoding object  as per OAS 3.0.3
type Encoding struct {
	ContentType   string            `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]Header `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string            `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool              `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool              `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}

//ExternalDocumentation object  as per https://spec.openapis.org/oas/v3.1.0#externalDocumentationObject
type ExternalDocumentation struct {
	URL         string `json:"url" yaml:"url,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

//Example object  as per OAS 3.0.3 https://spec.openapis.org/oas/v3.1.0#example-object
type Example struct {
	Reference
	Summary       string      `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   string      `json:"description,omitempty" yaml:"description,omitempty"`
	Value         interface{} `json:"example,omitempty" yaml:"example,omitempty"`
	ExternalValue string      `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
}

// Reference Type
type Reference struct {
	Ref         *string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string  `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
}

type Extension map[string]interface{}

//Schema Object for
type Schema struct {
	Reference
	ID                   string                `json:"id,omitempty" yaml:"id,omitempty"`
	Schema               string                `json:"-" yaml:"-"`
	Description          string                `json:"description,omitempty" yaml:"description,omitempty"`
	Type                 string                `json:"type,omitempty" yaml:"type,omitempty"`
	Nullable             bool                  `json:"nullable,omitempty" yaml:"nullable,omitempty"`
	Format               *string               `json:"format,omitempty" yaml:"format,omitempty"`
	Title                string                `json:"title,omitempty" yaml:"title,omitempty"`
	Default              interface{}           `json:"default,omitempty" yaml:"default,omitempty"`
	Maximum              *float64              `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum     *float64              `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum              *float64              `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum     *float64              `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength            *int                  `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength            *int                  `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern              *string               `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems             *int                  `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems             *int                  `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems          bool                  `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MultipleOf           *float64              `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Enum                 []interface{}         `json:"enum,omitempty" yaml:"enum,omitempty"`
	MaxProperties        *int                  `json:"maxProperties,omitempty" yaml:"maxProperties,omitempty"`
	MinProperties        *int                  `json:"minProperties,omitempty" yaml:"minProperties,omitempty"`
	Required             []string              `json:"required,omitempty" yaml:"required,omitempty"`
	Items                *Schema               `json:"items,omitempty" yaml:"items,omitempty"`
	AllOf                []*Schema             `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf                []*Schema             `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf                []*Schema             `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	Not                  *Schema               `json:"not,omitempty" yaml:"not,omitempty"`
	Properties           map[string]*Schema    `json:"properties,omitempty" yaml:"properties,omitempty"`
	AdditionalProperties interface{}           `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	AdditionalItems      *Schema               `json:"additionalItems,omitempty" yaml:"additionalItems,omitempty"`
	Xml                  *Xml                  `json:"xml,omitempty" yaml:"xml,omitempty"`
	ReadOnly             bool                  `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	WriteOnly            bool                  `json:"writeOnly,omitempty" yaml:"writeOnly,omitempty"`
	Discriminator        Discriminator         `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
	ExternalDocs         ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Example              interface{}           `json:"example,omitempty" yaml:"example,omitempty"`
	Examples             []interface{}         `json:"examples,omitempty" yaml:"examples,omitempty"`
	Deprecated           bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Extension
}

type Xml struct {
	Name      *string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Prefix    *string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Attribute *bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Wrapped   *bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
}

type PatternProperties map[string]*Schema

type Discriminator struct {
	PropertyName string            `json:"propertyName,omitempty" yaml:"wrapped,omitempty"`
	Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
}
