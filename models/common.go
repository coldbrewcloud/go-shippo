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

	ObjectStatusWaiting           = "WAITING"
	ObjectStatusQueued            = "QUEUED"
	ObjectStatusSuccess           = "SUCCESS"
	ObjectStatusError             = "ERROR"
	ObjectStatusRefunded          = "REFUNDED"
	ObjectStatusRefundPending     = "REFUNDPENDING"
	ObjectStatusRefundRejected    = "REFUNDREJECTED"
	ObjectStatusValidating        = "VALIDATING"
	ObjectStatusValid             = "VALID"
	ObjectStatusInvalid           = "INVALID"
	ObjectStatusIncomplete        = "INCOMPLETE"
	ObjectStatusPurchasing        = "PURCHASING"
	ObjectStatusPurchased         = "PURCHASED"
	ObjectStatusTransactionFailed = "TRANSACTION_FAILED"

	InsuranceProviderFedEx  = "FEDEX"
	InsuranceProviderUPS    = "UPS"
	InsuranceProviderOnTrac = "ONTRAC"

	CODPaymentMethodSecuredFunds = "SECURED_FUNDS"
	CODPaymentMethodCash         = "CASH"
	CODPaymentMethodAny          = "ANY"

	DistanceUnitCentiMeter = "cm"
	DistanceUnitInch       = "in"
	DistanceUnitFoot       = "ft"
	DistanceUnitMilliMeter = "mm"
	DistanceUnitMeter      = "m"
	DistanceUnitYard       = "yd"

	MassUnitGram     = "g"
	MassUnitOunce    = "oz"
	MassUnitPound    = "lb"
	MassUnitKiloGram = "kg"

	LabelFileTypePNG    = "PNG"
	LabelFileTypePDF    = "PDF"
	LabelFileTypePDF4X6 = "PDF_4X6"
	LabelFileTypeZPLII  = "ZPLII"
)

type CommonOutputFields struct {
	ObjectState   string    `json:"object_state,omitempty"`
	ObjectCreated time.Time `json:"object_created"`
	ObjectUpdated time.Time `json:"object_updated"`
	ObjectID      string    `json:"object_id,omitempty"`
	ObjectOwner   string    `json:"object_owner,omitempty"`
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
