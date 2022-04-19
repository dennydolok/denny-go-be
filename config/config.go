package config

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	SECRET_KEY  string
}

func InitConfig() Config {
	return Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "Xu33P90YUhaCua89GlS5",
		DB_PORT:     "5756",
		DB_NAME:     "railway",
		DB_HOST:     "containers-us-west-38.railway.app",
		SECRET_KEY:  "secret",
	}
}
