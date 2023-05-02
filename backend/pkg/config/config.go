package config

type appConfig struct {
	HTTP *HTTPConfig
	DB   *DBConfig
}

type HTTPConfig struct {
	Addr string
}

type DBConfig struct {
	Username string
	Password string
	Addr     string
	DBName   string
}

func LoadConfig() *appConfig {
	httpConfig := &HTTPConfig{
		Addr: ":8080",
	}

	dbConfig := &DBConfig{
		Username: "user",
		Password: "password",
		Addr:     "127.0.0.1:3306",
		DBName:   "articles",
	}

	return &appConfig{
		HTTP: httpConfig,
		DB:   dbConfig,
	}
}
