package config

import (
	"fmt"
	"os"
)

// Init the configs required in the package
func Init() {
	fmt.Println("Initializing Module")
	Load.RedisURL = os.Getenv("REDIS_URL")
	Load.PORT = os.Getenv("PORT")

	//  ensure that both the essential configs have a valid setting
	if Load.RedisURL == "" {
		fmt.Println("REDIS_URL required to run")
		os.Exit(1)
	}

	if Load.PORT == "" {
		fmt.Println("PORT required to run")
		os.Exit(2)
	}

}

var Load Config
