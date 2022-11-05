package config

import "fmt"

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
	DatabaseUser     string
}

var CONFIG *Config

func LoadConfig() {
	CONFIG = &Config{}
	// Service
	CONFIG.ServiceHost = "localhost"
	CONFIG.ServicePort = 8080

	// Db
	CONFIG.DatabaseHost = "localhost"
	CONFIG.DatabasePort = 6432
	CONFIG.DatabasePassword = "f49ac135aef919ca108a3e907d477493"
	CONFIG.DatabaseName = "dev"
	CONFIG.DatabaseUser = "dev"
	CONFIG.DatabaseUrl = fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s",
		CONFIG.DatabaseUser,
		CONFIG.DatabasePassword,
		CONFIG.DatabaseHost,
		CONFIG.DatabasePort,
		CONFIG.DatabaseName,
	)
}
