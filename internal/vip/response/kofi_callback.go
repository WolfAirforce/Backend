package response

import (
	"time"
)

type KofiCallbackData struct {
	MessageID                  string    `json:"message_id"`
	Timestamp                  time.Time `json:"timestamp"`
	Type                       string    `json:"type"`
	IsPublic                   bool      `json:"is_public"`
	FromName                   string    `json:"from_name"`
	Message                    string    `json:"message"`
	Amount                     float64   `json:"amount,string"`
	URL                        string    `json:"url"`
	Email                      string    `json:"email"`
	Currency                   string    `json:"currency"`
	IsSubscriptionPayment      bool      `json:"is_subscription_payment"`
	IsFirstSubscriptionPayment bool      `json:"is_first_subscription_payment"`
	KofiTransactionID          string    `json:"kofi_transaction_id"`
	VerificationToken          string    `json:"verification_token"`
}
