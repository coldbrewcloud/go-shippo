package models

import (
	"encoding/json"
	"time"
)

const (
	ObjectPurposeQuote    = "QUOTE"
	ObjectPurposePurchase = "PURCHASE"

	ObjectStateValid      = "VALID"
	ObjectStateInvalid    = "INVALID"
	ObjectStateIncomplete = "INCOMPLETE"

	ObjectStatusWaiting = "WAITING"
	ObjectStatusQueued  = "QUEUED"
	ObjectStatusSuccess = "SUCCESS"
	ObjectStatusError   = "ERROR"

	InsuranceProviderFedEx  = "FEDEX"
	InsuranceProviderUPS    = "UPS"
	InsuranceProviderOnTrac = "ONTRAC"

	CODPaymentMethodSecuredFunds = "SECURED_FUNDS"
	CODPaymentMethodCash         = "CASH"
	CODPaymentMethodAny          = "ANY"
)

type CommonOutputFields struct {
	ObjectState   string    `json:"object_state"`
	ObjectCreated time.Time `json:"object_created"`
	ObjectUpdated time.Time `json:"object_updated"`
	ObjectID      string    `json:"object_id"`
	ObjectOwner   string    `json:"object_owner"`
}

type OutputMessage struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type ListAPIOutput struct {
	Count           int               `json:"count"`
	NextPageURL     *string           `json:"next"`
	PreviousPageURL *string           `json:"previous"`
	Results         []json.RawMessage `json:"results"`
}

type COD struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
}
