package middlewares

import (
	"os"

	"github.com/gofiber/basicauth"
)

var cfg = basicauth.Config{
	Users: map[string]string{
		os.Getenv("USER"): os.Getenv("PASSWORD"),
	},
}

// GetConfig returns the basicauth configured to the app
func GetConfig() basicauth.Config {
	return cfg
}
