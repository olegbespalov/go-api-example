package config

import "os"

var (
	// AppPort port of the application
	AppPort string
)

func init() {
	AppPort = os.Getenv("APP_PORT")
}
