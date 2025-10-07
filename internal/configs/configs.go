package configs

import (
	"CSR/internal/models"
	"encoding/json"
	"fmt"

	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	
)

var (
	AppSettings models.Config
)

func GetConfig(path string) error {
	
	logger:=zerolog.New(os.Stdin).With().Timestamp().Logger()
	err := godotenv.Load("internal/.env")
	if err != nil {
		logger.Err(err).Msg("error while getting .env")
		return fmt.Errorf("error while getting .env, %v", err)
	}

	f, err := os.Open(path)
	if err != nil {
		logger.Err(err).Msg("error while opening file config.json")
		return fmt.Errorf("error while opening file config.json\nerr:%v\n", err)
	}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&AppSettings)
	if err != nil {
		logger.Err(err).Msg("error while desconding config file")
		return fmt.Errorf("error while decoding  file config.json\nerr:%v\n", err)

	}
	
	return nil
}
