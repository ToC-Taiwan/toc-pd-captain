package config

type Database struct {
	DBName  string `env:"DB_NAME" env-required:"true"`
	URL     string `env:"DB_URL" env-required:"true"`
	PoolMax int    `env:"DB_POOL_MAX" env-required:"true"`
}

type Server struct {
	HTTP                      string `env:"HTTP" env-required:"true"`
	RouterDebugMode           string `env:"GIN_MODE" env-required:"true"`
	DisableSwaggerHTTPHandler string `env:"DISABLE_SWAGGER_HTTP_HANDLER" env-required:"true"`
}
