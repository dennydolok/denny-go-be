package config

import "os"

type Config struct {
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	SECRET_KEY     string
	SERVER_ADDRESS string
}

func InitConfig() Config {
	return Config{
		DB_USERNAME:    GetValue("DB_USERNAME", "root"),
		DB_PASSWORD:    GetValue("DB_PASSWORD", "root"),
		DB_PORT:        GetValue("DB_PORT", "3306"),
		DB_NAME:        GetValue("DB_NAME", "db_be"),
		DB_HOST:        GetValue("DB_HOST", "0.0.0.0"),
		SECRET_KEY:     GetValue("JWT_KEY", "secret"),
		SERVER_ADDRESS: GetValue("SERVER_ADDRESS", "0.0.0.0:8080"),
	}
}

func GetValue(envKey, defaultKey string) string {
	if val, exist := os.LookupEnv(envKey); exist {
		return val
	}
	return defaultKey
}
