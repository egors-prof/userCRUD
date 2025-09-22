package configs

import (
	"CSR/Internal/models"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	AppSettings models.Config
)

func GetConfig(path string) error {

	err := godotenv.Load("internal/.env")
	if err != nil {
		return fmt.Errorf("error while getting .env, %v", err)
	}

	f, err := os.Open(path)
	if err != nil {

		return fmt.Errorf("error while opening file config.json\nerr:%v\n", err)
	}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&AppSettings)
	if err != nil {
		return fmt.Errorf("error while decoding  file config.json\nerr:%v\n", err)

	}
	return nil
}
