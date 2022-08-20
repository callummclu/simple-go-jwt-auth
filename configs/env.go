package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func getEnvByName(name string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	return os.Getenv(name), nil

}

func EnvPORT() (string, error) {
	port, err := getEnvByName("SERVICE_PORT")
	return ":" + port, err
}
