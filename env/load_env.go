package env

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvFileNoError loads the passed env file.
// If there is an error it is logged without stopping application.
func LoadEnvFileNoError(envFilePath string) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Printf("error loading env file at: %s, err: %s", envFilePath, err)
	}
}

// LoadEnvFileWithError loads the passed env file.
// If there is an error it is returned.
func LoadEnvFileWithError(envFilePath string) error {
	return godotenv.Load(envFilePath)
}
