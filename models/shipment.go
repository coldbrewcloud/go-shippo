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
	AddressFrom        interface{}    `json:"address_from"`                  // string or *AddressInput
	AddressTo          interface{}    `json:"address_to"`                    // string or *AddressInput
	Parcels            interface{}    `json:"parcels"`                       // string, []string, *ParcelInput, or []*ParcelInput
	AddressReturn      interface{}    `json:"address_return,omitempty"`      // string or *AddressInput
	CustomsDeclaration interface{}    `json:"customs_declaration,omitempty"` // string or *CustomsDeclarationInput
	ShipmentDate       time.Time      `json:"shipment_date,omitempty"`
	CarrierAccounts    []string       `json:"carrier_accounts,omitempty"`
	Metadata           string         `json:"metadata,omitempty"`
	Extra              *ShipmentExtra `json:"extra,omitempty"`
	Async              bool           `json:"async"`
}

type ShipmentExtra struct {
	SignatureConfirmation   string             `json:"signature_confirmation,omitempty"`
	SaturdayDelivery        bool               `json:"saturday_delivery"`
	BypassAddressValidation bool               `json:"bypass_address_validation"`
	RequestRetailRates      bool               `json:"request_retail_rates"`
	CustomerBranch          string             `json:"customer_branch,omitempty"`
	Billing                 *ShipmentBilling   `json:"billing,omitempty"`
	COD                     *ShipmentCOD       `json:"COD,omitempty"`
	Insurance               *ShipmentInsurance `json:"insurance,omitempty"`
	USPSSortType            string             `json:"usps_sort_type,omitempty"`
	USPSEntryFacility       string             `json:"usps_entry_facility,omitempty"`
	DangerousGoodsCode      string             `json:"dangerous_goods_code,omitempty"`
	IsReturn                bool               `json:"is_return,omitempty"`
	Reference1              string             `json:"reference_1,omitempty"`
	Reference2              string             `json:"reference_2,omitempty"`
}

type ShipmentCOD struct {
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
}

type ShipmentInsurance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Provider string `json:"provider,omitempty"`
	Content  string `json:"content"`
}

type ShipmentBilling struct {
	Type    string `json:"type"`
	Account string `json:"account"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

// See https://goshippo.com/docs/reference#shipments
type Shipment struct {
	CommonOutputFields
	Status             string              `json:"status,omitempty"`
	AddressFrom        *Address            `json:"address_from,omitempty"`
	AddressTo          *Address            `json:"address_to,omitempty"`
	Parcels            []*Parcel           `json:"parcels,omitempty"`
	AddressReturn      *Address            `json:"address_return,omitempty"`
	CustomsDeclaration *CustomsDeclaration `json:"customs_declaration,omitempty"`
	Rates              []*Rate             `json:"rates,omitempty"`
	Messages           []*OutputMessage    `json:"messages,omitempty"`
	Test               bool                `json:"test,omitempty"`
}
