package config

import "os"

var (
	// AppPort port of the application
	AppPort string

	// AppName is an application name
	AppName string
)

func init() {
	AppPort = os.Getenv("APP_PORT")
	AppName = os.Getenv("APP_NAME")
}
