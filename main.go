package main

import (
    "fmt"
    "os"
    "proj-mido/stripe-gateway/Config"
    "proj-mido/stripe-gateway/Routes"
    "github.com/gin-contrib/cors"
    "github.com/joho/godotenv"

)


func main() {
    // Open the connection
        
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
    config := Config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	Config.InitDB(config)
       // Setup the router
    r := Routes.SetupRouter()

    // Middleware
    r.Use(cors.Default())

    // Run the server
    r.Run()
}
