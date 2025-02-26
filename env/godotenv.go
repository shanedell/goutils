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

// LoadEnvFileNoErrorNoLog loads the passed env file.
// If there is an error nothing happens.
func LoadEnvFileNoErrorNoLog(envFilePath string) {
	_ = godotenv.Load(envFilePath)
}

// LoadEnvFileWithError loads the passed env file.
// If there is an error it is returned.
func LoadEnvFileWithError(envFilePath string) error {
	return godotenv.Load(envFilePath)
}
