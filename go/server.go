package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

type PaymentParams struct {
	AmountUSD int64 `json:"amount_usd" validate:"numeric"`
}

func JSON[T any](w http.ResponseWriter, status int, data T) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ReadParams[T any](body io.ReadCloser) (*T, error) {
	var params T
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := json.NewDecoder(body).Decode(&params)
	if err != nil {
		return new(T), err
	}

	err = validate.Struct(params)
	if err != nil {
		return new(T), err
	}
	return &params, nil
}
