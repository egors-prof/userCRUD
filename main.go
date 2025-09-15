package main

import (
	"CSR/Internal/Controller"
	"CSR/Internal/Repository"
	"CSR/Internal/Service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sqlx.Open("postgres", "host=localhost user=postgres password=12345 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connected to postgres")

	repository := Repository.NewRepository(db)
	service := Service.NewService(repository)
	controller := Controller.NewController(service)

	if err = controller.RunServer(":8888"); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}

}
