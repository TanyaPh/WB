package entity

type Payment struct {
	Id           int
	Transaction  string
	RequestID    string
	Currency     string
	Provider     string
	Amount       int
	PaymentDT    int64
	Bank         string
	DeliveryCost int
	GoodTotal    int
	CustomFree   int
}
