package models

type Config struct {
	DbName string
	DbPort string
	DbPass string
	DbHost string
	DbUser string
}

func NewConfig(dbHost, dbPort, dbUser, dbPass, dbName string) *Config {
	return &Config{
		DbName: dbName,
		DbPort: dbPort,
		DbPass: dbPass,
		DbHost: dbHost,
		DbUser: dbUser,
	}
}
