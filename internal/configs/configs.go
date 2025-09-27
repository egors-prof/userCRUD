package configs

import (
	"CSR/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AppSettings models.Config
)

func GetConfig(path string) error {

	err := godotenv.Load("Internal/.env")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("error while getting .env, %v", err)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("error while opening file config.json\nerr:%v\n", err)
	}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&AppSettings)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("error while decoding  file config.json\nerr:%v\n", err)

	}
	return nil
}
