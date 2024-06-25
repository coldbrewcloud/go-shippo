package models

const (
	CustomsNonDeliveryOptionAbandon = "ABANDON"
	CustomsNonDeliveryOptionReturn  = "RETURN"

	CustomsContentsTypeDocuments            = "DOCUMENTS"
	CustomsContentsTypeGift                 = "GIFT"
	CustomsContentsTypeSample               = "SAMPLE"
	CustomsContentsTypeHumanitarianDonation = "HUMANITARIAN_DONATION"
	CustomsContentsTypeMerchandise          = "MERCHANDISE"
	CustomsContentsTypeReturnMerchandise    = "RETURN_MERCHANDISE"
	CustomsContentsTypeOther                = "OTHER"

	CustomsEEL_PFC_NOEEI_30_37_A = "NOEEI_30_37_a"
	CustomsEEL_PFC_NOEEI_30_37_H = "NOEEI_30_37_h"
	CustomsEEL_PFC_NOEEI_30_37_F = "NOEEI_30_37_f"
	CustomsEEL_PFC_NOEEI_30_36   = "NOEEI_30_36"
	CustomsEEL_PFC_AES_ITN       = "AES_ITN"

	CustomsIncotermDDP = "DDP"
	CustomsIncotermDDU = "DDU"
	CustomsIncotermCPT = "CPT"
	CustomsIncotermCIP = "CIP"
)

// See https://goshippo.com/docs/reference#customsitems
type CustomsItemInput struct {
	Description   string `json:"description"`
	Quantity      int    `json:"quantity"`
	NetWeight     string `json:"net_weight"`
	MassUnit      string `json:"mass_unit"`
	ValueAmount   string `json:"value_amount"`
	ValueCurrency string `json:"value_currency"`
	OriginCountry string `json:"origin_country"`
	TariffNumber  string `json:"tariff_number,omitempty"`
	Metadata      string `json:"metadata,omitempty"`
}

// See https://goshippo.com/docs/reference#customsitems
type CustomsItem struct {
	CustomsItemInput
	CommonOutputFields
	ObjectState string `json:"object_state,omitempty"`
	Test        bool   `json:"test"`
}

// See https://goshippo.com/docs/reference#customsdeclarations
type CustomsDeclarationInput struct {
	CertifySigner       string      `json:"certify_signer"`
	Certify             bool        `json:"certify"`
	Items               interface{} `json:"items"` // []string or []*CustomsItemInput
	NonDeliveryOption   string      `json:"non_delivery_option"`
	ContentsType        string      `json:"contents_type"`
	ContentsExplanation string      `json:"contents_explanation,omitempty"`
	ExporterReference   string      `json:"exporter_reference,omitempty"`
	importerReference   string      `json:"importer_reference,omitempty"`
	Invoice             string      `json:"invoice,omitempty"`
	License             string      `json:"license,omitempty"`
	Certificate         string      `json:"certificate,omitempty"`
	Notes               string      `json:"notes,omitempty"`
	EEL_PFC             string      `json:"eel_pfc,omitempty"`
	AES_ITN             string      `json:"aes_itn,omitempty"`
	Incoterm            string      `json:"incoterm,omitempty"`
	Metadata            string      `json:"metadata,omitempty"`
	Disclaimer          string      `json:"disclaimer,omitempty"`
}

// See https://goshippo.com/docs/reference#customsdeclarations
type CustomsDeclaration struct {
	CustomsDeclarationInput
	CommonOutputFields
	ObjectState string `json:"object_state,omitempty"`
	Test        bool   `json:"test"`
}
