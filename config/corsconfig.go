package config

import (
	"github.com/gin-contrib/cors"
	"time"
)

func CorsConfig() cors.Config {
	c := cors.Config{
		AllowOrigins: []string{
			"https://joe-take.github.io",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"OPTIONS",
		},

		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},

		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}

	return c
}
