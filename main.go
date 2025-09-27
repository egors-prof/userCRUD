package main

import (
	"CSR/internal/configs"
	"CSR/internal/controller"
	"CSR/internal/repository"
	"CSR/internal/service"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// @title onlineshop
// @contact.name onlineshop api
// @contact.url http://test.com
// @contact.email test@test.com
func main() {

	err := configs.GetConfig("Internal/configs/configs.json")
	if err != nil {
		log.Fatal("error getting configs")
	}
	fmt.Println(configs.AppSettings)
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
		log.Fatal(err)
	}

	log.Println("successfully connected to postgres")
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.AppSettings.RedisParams.Address,
		Password: os.Getenv("REDIS_PASS"),
		DB:       configs.AppSettings.RedisParams.DB,
	})
	cache := repository.NewCache(rdb)
	repository := repository.NewRepository(db, cache)
	service := service.NewService(repository, cache)
	controller := controller.NewController(service)

	if err = controller.RunServer(); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

}
