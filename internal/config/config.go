package config

type Config struct {
	// Service
	ServiceHost string
	ServicePort int
	// Db
	DatabaseUrl      string
	DatabaseHost     string
	DatabasePort     int
	DatabasePassword string
	DatabaseName     string
}

var CONFIG *Config

func LoadConfig() {
	CONFIG = &Config{}
	// Service
	CONFIG.ServiceHost = "localhost"
	CONFIG.ServicePort = 8080

	// Db
	CONFIG.DatabaseUrl = "postgres://{username}:{password}@{host}:{port}/{db_name}"
	CONFIG.DatabaseHost = ""
	CONFIG.DatabasePort = 5672
	CONFIG.DatabasePassword = ""
	CONFIG.DatabaseName = ""
}
