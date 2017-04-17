package models

import "time"

const (
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
	AddressFrom        string    `json:"address_from"`
	AddressTo          string    `json:"address_to"`
	Parcels            []string  `json:"parcels"`
	ShipmentDate       time.Time `json:"shipment_date,omitempty"`
	AddressReturn      string    `json:"address_return,omitempty"`
	CustomsDeclaration string    `json:"customs_declaration,omitempty"`
	CarrierAccounts    []string  `json:"carrier_accounts,omitempty"`
	Metadata           string    `json:"metadata,omitempty"`
	Async              bool      `json:"async,omitempty"`
	Extra              *Extra    `json:"extra,omitempty"`
}

// Represents shipment response back from shippo
type ShipmentOutput struct {
	AddressFrom        *Address  `json:"address_from"`
	AddressTo          *Address  `json:"address_to"`
	Parcels            []*Parcel `json:"parcels"`
	ShipmentDate       time.Time `json:"shipment_date,omitempty"`
	AddressReturn      *Address  `json:"address_return,omitempty"`
	CustomsDeclaration string    `json:"customs_declaration,omitempty"`
	CarrierAccounts    []string  `json:"carrier_accounts,omitempty"`
	Metadata           string    `json:"metadata,omitempty"`
	Async              bool      `json:"async,omitempty"`
	Extra              *Extra    `json:"extra,omitempty"`
}

type Extra struct {
	SignatureConfirmation   string           `json:"signature_confirmation,omitempty"`
	SaturdayDelivery        bool             `json:"saturday_delivery"`
	Insurance               *Insurance       `json:"insurance"`
	BypassAddressValidation bool             `json:"bypass_address_validation"`
	UseManifests            bool             `json:"use_manifests"`
	RequestRetailRates      bool             `json:"request_retail_rates"`
	Billing                 *ShipmentBilling `json:"billing,omitempty"`
	COD                     *COD             `json:"COD,omitempty"`
	USPSSortType            string           `json:"usps_sort_type,omitempty"`
	USPSEntryFacility       string           `json:"usps_entry_facility,omitempty"`
	DangerousGoodsCode      string           `json:"dangerous_goods_code,omitempty"`
	IsReturn                bool             `json:"is_return,omitempty"`
	Reference1              string           `json:"reference_1,omitempty"`
	Reference2              string           `json:"reference_2,omitempty"`
}

type Insurance struct {
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Content  string `json:"content,omitempty"`
	Provider string `json:"provider,omitempty"`
}

type ShipmentBilling struct {
	Type    string `json:"type"`
	Account string `json:"account"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

// See https://goshippo.com/docs/reference#shipments
type Shipment struct {
	ShipmentOutput
	CommonOutputFields
	Status   string           `json:"status"`
	Rates    []*Rate          `json:"rates"`
	Messages []*OutputMessage `json:"messages"`
	Extra    *Extra           `json:"extra"`
	Test     bool             `json:"test"`
}
