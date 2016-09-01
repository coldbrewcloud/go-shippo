package models

import "time"

const (
	ShipmentSubmissionTypeDropOff = "DROPOFF"
	ShipmentSubmissionTypePickUp  = "PICKUP"

	ShipmentExtraSignatureConfirmationStandard  = "STANDARD"
	ShipmentExtraSignatureConfirmationAdult     = "ADULT"
	ShipmentExtraSignatureConfirmationCertified = "CERTIFIED"

	ShipmentBillingTypeSender     = "SENDER"
	ShipmentBillingTypeRecipient  = "RECIPIENT"
	ShipmentBillingTypeThirdParty = "THIRD_PARTY"

	ShipmentUSPSSortTypeNDC          = "NDC"
	ShipmentUSPSSortTypeFiveDigit    = "FiveDigit"
	ShipmentUSPSSortTypeMixedNDC     = "MixedNDC"
	ShipmentUSPSSortTypeNonPresorted = "Nonpresorted"
	ShipmentUSPSSortTypePresorted    = "Presorted"
	ShipmentUSPSSortTypeSCF          = "SCF"
	ShipmentUSPSSortTypeSinglePiece  = "SinglePiece"
	ShipmentUSPSSortTypeThreeDigit   = "ThreeDigit"

	ShipmentUSPSEntryFacilityDNDC  = "DNDC"
	ShipmentUSPSEntryFacilityDDU   = "DDU"
	ShipmentUSPSEntryFacilityDSCF  = "DSCF"
	ShipmentUSPSEntryFacilityONDC  = "ONDC"
	ShipmentUSPSEntryFacilityOther = "Other"

	ShipmentDangerousGoodsCode01 = "01"
	ShipmentDangerousGoodsCode02 = "02"
	ShipmentDangerousGoodsCode03 = "03"
	ShipmentDangerousGoodsCode04 = "04"
	ShipmentDangerousGoodsCode05 = "05"
	ShipmentDangerousGoodsCode06 = "06"
	ShipmentDangerousGoodsCode07 = "07"
	ShipmentDangerousGoodsCode08 = "08"
	ShipmentDangerousGoodsCode09 = "019"
)

// See https://goshippo.com/docs/reference#shipments
type ShipmentInput struct {
	ObjectPurpose      string         `json:"object_purpose"`
	AddressFrom        string         `json:"address_from"`
	AddressTo          string         `json:"address_to"`
	Parcel             string         `json:"parcel"`
	ReturnOf           string         `json:"return_of,omitempty"`
	SubmissionType     string         `json:"submission_type,omitempty"`
	SubmissionDate     time.Time      `json:"submission_date,omitempty"`
	AddressReturn      string         `json:"address_return,omitempty"`
	CustomsDeclaration string         `json:"customs_declaration,omitempty"`
	InsuranceAmount    string         `json:"insurance_amount,omitempty"`
	InsuranceCurrency  string         `json:"insurance_currency,omitempty"`
	Reference1         string         `json:"reference_1,omitempty"`
	Reference2         string         `json:"reference_2,omitempty"`
	CarrierAccounts    []string       `json:"carrier_accounts,omitempty"`
	Metadata           string         `json:"metadata,omitempty"`
	Async              bool           `json:"async"`
	Extra              *ShipmentExtra `json:"extra,omitempty"`
}

type ShipmentExtra struct {
	SignatureConfirmation   string           `json:"signature_confirmation,omitempty"`
	SaturdayDelivery        bool             `json:"saturday_delivery"`
	InsuranceContent        string           `json:"insurance_content,omitempty"`
	InsuranceProvider       string           `json:"insurance_provider,omitempty"`
	BypassAddressValidation bool             `json:"bypass_address_validation"`
	UseManifests            bool             `json:"use_manifests"`
	RequestRetailRates      bool             `json:"request_retail_rates"`
	Billing                 *ShipmentBilling `json:"billing,omitempty"`
	COD                     *COD             `json:"COD,omitempty"`
	USPSSortType            string           `json:"usps_sort_type,omitempty"`
	USPSEntryFacility       string           `json:"usps_entry_facility,omitempty"`
	DangerousGoodsCode      string           `json:"dangerous_goods_code,omitempty"`
}

type ShipmentBilling struct {
	Type    string `json:"type"`
	Account string `json:"account"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

// See https://goshippo.com/docs/reference#shipments
type Shipment struct {
	ShipmentInput
	CommonOutputFields
	ObjectStatus string           `json:"object_status"`
	RatesURL     string           `json:"rates_url"`
	RatesList    []*Rate          `json:"rates_list"`
	Messages     []*OutputMessage `json:"messages"`
}
