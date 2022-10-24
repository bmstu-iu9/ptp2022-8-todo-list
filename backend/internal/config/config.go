package config

import "os"

var config = map[string]string{
	"HTTP_PORT":    "8080",
	"HTTP_HOST":    "0.0.0.0",
	"API_SERVER":   "https://ptp.starovoytovai.ru/api/v1",
	"DB_PORT":      "5432",
	"DB_HOST":      "localhost",
	"DB_USER":      "slava",
	"DB_NAME":      "slavatidika",
	"DB_PASSWORD":  "asdwasd4545",
	"DB_SSL_MODE":  "disable",
	"RUNTIME_MODE": "debug",
	"EMAIL_HOST":      "smtp.gmail.com",
	"EMAIL_SENDER":    "slavatidika@gmail.com",
	"EMAIL_PASSWORD":  "ojlakqwiiuvknvcx",
}

func init() {
	for variable := range config {
		if envValue := os.Getenv(variable); envValue != "" {
			config[variable] = envValue
		}
	}
}

// Get returns configuration parameter named variable.
func Get(variable string) string {
	return config[variable]
}
