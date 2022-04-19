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
		DB_PASSWORD: "",
		DB_PORT:     "3306",
		DB_NAME:     "training_clean",
		DB_HOST:     "localhost",
		SECRET_KEY:  "secret",
	}
}
