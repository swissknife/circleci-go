package shared

type SchemeAPIKeyHeader struct {
	APIKey string `security:"name=Circle-Token"`
}

type SchemeBasicAuth struct {
	Password string `security:"name=password"`
	Username string `security:"name=username"`
}

type SchemeAPIKeyQuery struct {
	APIKey string `security:"name=circle-token"`
}

type Security struct {
	APIKeyHeader *SchemeAPIKeyHeader `security:"scheme,type=apiKey,subtype=header"`
	BasicAuth    *SchemeBasicAuth    `security:"scheme,type=http,subtype=basic"`
	APIKeyQuery  *SchemeAPIKeyQuery  `security:"scheme,type=apiKey,subtype=query"`
}
