package api

import (
	"fmt"
	"net/http"
)

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /payments", cors(handlePayments))
	mux.HandleFunc("POST /subscriptions", cors(handleSubscriptions))
	mux.HandleFunc("POST /checkout-session", cors(handleCheckoutSession))
	mux.HandleFunc("GET /health", cors(handleHealthCheck))
}

func cors(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		fmt.Printf("headers: %+v", w.Header())
		handler(w, r)
	}
}
