package config

type AuthConfig struct {
	TokenSigningKey        string
	TokenExpirationSeconds string
}

type DatabaseConfig struct {
	UserName  string
	Password  string
	Host      string
	Port      string
	Schema    string
	Charset   string
	ParseTime string
	Loc       string
}

type ServerConfig struct {
	Port string
}
