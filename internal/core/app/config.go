package app

type DatabaseConfig struct {
	Sslmode        string
	ConnectTimeout int
	Host           string
	Port           int
	Name           string
	User           string
	Password       string
}

type CacheConfig struct {
	Type  string
	Redis RedisConfig
}

type RedisConfig struct {
	Addr string
	DB   int
}

type Config struct {
	Deployment  string
	Credentials CredentialsConfig
	Server      ServerConfig
	Database    DatabaseConfig
	Cache       CacheConfig
}

type CredentialsConfig struct {
	Secret string
	Key    string
}

type ServerConfig struct {
	Port int
	Addr string
}
