package models

import "time"

type OrderInput struct {
	ToAddress            interface{} `json:"to_address"`
	LineItems            []LineItem  `json:"line_items"`
	PlacedAt             time.Time   `json:"placed_at"`
	OrderNumber          string      `json:"order_number"`
	OrderStatus          string      `json:"order_status"`
	ShippingCost         string      `json:"shipping_cost"`
	ShippingCostCurrency string      `json:"shipping_cost_currency"`
	ShippingMethod       string      `json:"shipping_method"`
	SubtotalPrice        string      `json:"subtotal_price"`
	TotalPrice           string      `json:"total_price"`
	TotalTax             string      `json:"total_tax"`
	Currency             string      `json:"currency"`
	Weight               string      `json:"weight"`
	WeightUnit           string      `json:"weight_unit"`
}

type LineItem struct {
	ObjectID           string    `json:"object_id"`
	Title              string    `json:"title"`
	VariantTitle       string    `json:"variant_title"`
	Sku                string    `json:"sku"`
	Quantity           int       `json:"quantity"`
	TotalPrice         string    `json:"total_price"`
	Currency           string    `json:"currency"`
	Weight             string    `json:"weight"`
	WeightUnit         string    `json:"weight_unit"`
	ManufactureCountry string    `json:"manufacture_country"`
	MaxShipTime        time.Time `json:"max_ship_time"`
	MaxDeliveryTime    time.Time `json:"max_delivery_time"`
}

type OrderResponse struct {
	ObjectID             string        `json:"object_id"`
	ObjectOwner          string        `json:"object_owner"`
	OrderNumber          string        `json:"order_number"`
	OrderStatus          string        `json:"order_status"`
	PlacedAt             time.Time     `json:"placed_at"`
	ToAddress            Address       `json:"to_address"`
	FromAddress          Address       `json:"from_address"` // Given it's null in the sample, this can be another struct or a pointer to a struct
	LineItems            []LineItem    `json:"line_items"`
	Items                []interface{} `json:"items"`
	Hidden               bool          `json:"hidden"`
	ShippingCost         string        `json:"shipping_cost"`
	ShippingCostCurrency string        `json:"shipping_cost_currency"`
	ShippingMethod       string        `json:"shipping_method"`
	ShopApp              string        `json:"shop_app"`
	SubtotalPrice        string        `json:"subtotal_price"`
	TotalPrice           string        `json:"total_price"`
	TotalTax             string        `json:"total_tax"`
	Currency             string        `json:"currency"`
	Transactions         []Transaction `json:"transactions"`
	Weight               string        `json:"weight"`
	WeightUnit           string        `json:"weight_unit"`
	Notes                interface{}   `json:"notes"` // Could be another struct or type based on actual data
}
