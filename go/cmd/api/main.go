package main

import (
	"errors"
	"fmt"
	"github.com/anmho/buy-me-a-boba/api"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
	"log"
	"net/http"
)

type Config struct {
	StripeKey string `env:"STRIPE_KEY"`
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("config: %+v\n", config)

	stripe.Key = config.StripeKey
	mux := api.MakeServer()

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("error serving: ", err)
		} else {
			fmt.Println("server closed")
		}
	}
}
