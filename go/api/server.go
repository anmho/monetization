package api

import (
	"fmt"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/stripe/stripe-go/v79/product"
	"github.com/stripe/stripe-go/v79/subscription"
	"log"
	"log/slog"
	"net/http"
)

type PaymentParams struct {
	AmountUSD       int64  `json:"amount_usd" validate:"required,numeric"`
	CustomerID      string `json:"customer_id" validate:"required"`
	PaymentMethodID string `json:"payment_method_id" validate:"required"`
}

type SubscriptionTier int

const (
	UnsetTier SubscriptionTier = iota
	NormalTier
	BudgetTier
	PremiumTier
)

func (t SubscriptionTier) String() string {
	switch t {
	case NormalTier:
		return "normal"
	case BudgetTier:
		return "budget"
	case PremiumTier:
		return "premium"
	default:
		return "unset"
	}
}

type SubscriptionParams struct {
	CustomerID            string `json:"customer_id" validate:"required"`
	SubscriptionProductID string `json:"subscription_product_id" validate:"required"`
	PaymentMethodID       string `json:"payment_method_id" validate:"required"`
}

type CheckoutItem struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int64  `json:"quantity" validate:"required"`
}

type CheckoutParams struct {
	CustomerID string         `json:"customer_id" validate:"required"`
	Items      []CheckoutItem `json:"items" validate:"required,dive"`
}

func MakeServer() http.Handler {
	mux := http.NewServeMux()
	registerRoutes(mux)
	return mux
}

// handleCreatePaymentIntent handles creating a payment for a custom payment flow.
func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	params, err := ReadParams[PaymentParams](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	intentParams := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(params.AmountUSD),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Customer:      stripe.String(params.CustomerID),
		PaymentMethod: stripe.String(params.PaymentMethodID),
	}

	payment, err := paymentintent.New(intentParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("payment created", slog.Any("payment", payment))
	resp := map[string]any{
		"client_secret": payment.ClientSecret,
	}
	JSON(w, http.StatusCreated, resp)
}

// handleSubscriptions handles subscribing a user for a customer subscription flow
func handleSubscriptions(w http.ResponseWriter, r *http.Request) {
	params, err := ReadParams[SubscriptionParams](r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n", params)
	sub, err := product.Get(params.SubscriptionProductID, &stripe.ProductParams{})
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	createSubParams := &stripe.SubscriptionParams{
		Customer: stripe.String(params.CustomerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(sub.DefaultPrice.ID),
			},
		},
		DefaultPaymentMethod: stripe.String(params.PaymentMethodID),
	}
	result, err := subscription.New(createSubParams)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	JSON(w, http.StatusCreated, result)
}

// handleCheckoutSession handles checking out a user when we want to use Stripe's pre-built payments page
func handleCheckoutSession(w http.ResponseWriter, r *http.Request) {
	params, err := ReadParams[CheckoutParams](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var lineItems []*stripe.CheckoutSessionLineItemParams
	//lineItems := make([]*stripe.CheckoutSessionLineItemParams, len(params.Items))
	for _, item := range params.Items {
		prod, err := product.Get(item.ProductID, nil)
		if err != nil {
			log.Println(err)
			return
		}
		slog.Info("item", slog.Any("item", item))
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(prod.DefaultPrice.ID),
			Quantity: stripe.Int64(item.Quantity),
		})
	}
	slog.Info("creating checkout session", slog.Any("lineItems", lineItems[0]))

	checkoutParams := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("http://localhost:3000/payments/success"),
		Customer:   stripe.String(params.CustomerID),
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems:  lineItems,
	}

	checkoutSession, err := session.New(checkoutParams)
	if err != nil {
		JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	JSON(w, http.StatusCreated, map[string]any{
		"checkout_url": checkoutSession.URL,
	})
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	JSON(w, http.StatusOK, map[string]any{
		"message": "hello",
	})
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {

}
