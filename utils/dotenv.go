package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// This function is used to get the proxy password from the .env file.
// input: path to .env file, key to get from .env file
// output: value of the key
func GoDotEnvVariable(dotenvpath string, key string) string {
	/* Load .env file and return the value of the key */

	err := godotenv.Load(dotenvpath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)

}
