package entity

type Payment struct {
	Id           int    `json:"-" db:"-"`
	Transaction  string `json:"transaction" db:"transaction"`
	RequestID    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int    `json:"amount" db:"amount"`
	PaymentDT    int64  `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost"`
	GoodTotal    int    `json:"goods_total" db:"goods_total"`
	CustomFree   int    `json:"custom_fee" db:"custom_fee"`
}
