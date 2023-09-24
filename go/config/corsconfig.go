package config

import (
	"github.com/gin-contrib/cors"
	"time"
)

func CorsConfig() cors.Config {
	c := cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
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
