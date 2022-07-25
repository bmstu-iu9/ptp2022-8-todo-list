package config

import "os"

// Port on which api server will run.
var Port = "8080"
// Network address on whitch api server will run.
var Host = "0.0.0.0"

var Server = "https://ptp.starovoytovai.ru/api/v1"

func init() {
	if port := os.Getenv("HTTP_PORT"); port != "" {
		Port = port
	}
	if host := os.Getenv("HTTP_HOST"); host != "" {
		Host = host
	}
}
