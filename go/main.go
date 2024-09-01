package main

import (
	"errors"
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"log"
	"net/http"
)

type Config struct {
	StripeKey string `env:"STRIPE_KEY"`
}

func MakeServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /payments", func(w http.ResponseWriter, r *http.Request) {
		params, err := ReadParams[PaymentParams](r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		intentParams := &stripe.PaymentIntentParams{
			Amount:       stripe.Int64(params.AmountUSD),
			Currency:     stripe.String(string(stripe.CurrencyUSD)),
			UseStripeSDK: stripe.Bool(true),
		}
		paymentIntent, err := paymentintent.New(intentParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("making payment")
		JSON(w, http.StatusCreated, paymentIntent)
	})

	mux.HandleFunc("POST /subscriptions", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("subscribing")
	})
	return mux
}

func main() {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		log.Fatalln(err)
	}

	stripe.Key = config.StripeKey

	mux := MakeServer()

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
