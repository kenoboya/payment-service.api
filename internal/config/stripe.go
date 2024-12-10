package config

import "github.com/stripe/stripe-go/v81"

type StripeConfig struct {
	APIKey        string
}

func (api *StripeConfig) Init() {
	stripe.Key = *stripe.String(api.APIKey)
}