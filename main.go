package main

import (
	"CSR/Internal/Controller"
	"CSR/Internal/Repository"
	"CSR/Internal/Service"
	"CSR/Internal/configs"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

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

	repository := Repository.NewRepository(db)
	service := Service.NewService(repository)
	controller := Controller.NewController(service)

	if err = controller.RunServer(); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

}
