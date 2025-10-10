package main

import (
	"CSR/internal/configs"
	"CSR/internal/controller"
	"CSR/internal/repository"
	"CSR/internal/service"
	"context"
	"fmt"

	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var ctx = context.Background()

// @title onlineshop
// @contact.name onlineshop api
// @contact.url http://test.com
// @contact.email test@test.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger := zerolog.New(os.Stdin).With().Timestamp().Logger()
	err := configs.GetConfig("internal/configs/configs.json")
	if err != nil {
		logger.Err(err).Msg("config error")
		return
	}
	postgresOpen := fmt.Sprintf(
		`host=%s
			user=%s
			password=%s
			dbname=%s
			sslmode=disable`,
		configs.AppSettings.PostgresParam.Host,
		configs.AppSettings.PostgresParam.User,
		os.Getenv("POSTGRES_PASS"),
		configs.AppSettings.PostgresParam.Database,
	)
	db, err := sqlx.Open("postgres", postgresOpen)
	if err != nil {
		logger.Err(err).Msg("error connecting to postgres")
		return
	}

	logger.Info().Msg("successfully connected to postgres")
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.AppSettings.RedisParams.Address,
		Password: os.Getenv("REDIS_PASS"),
		DB:       configs.AppSettings.RedisParams.DB,
	})

	cache := repository.NewCache(rdb, logger)
	repository := repository.NewRepository(db, cache, logger)
	service := service.NewService(repository, cache)
	controller := controller.NewController(service)

	if err = controller.RunServer(); err != nil {
		logger.Error().Err(err).Msg("error while trying to run server")
	}

	if err = db.Close(); err != nil {
		logger.Error().Err(err).Msg("error closing database")
		return
	}

}
