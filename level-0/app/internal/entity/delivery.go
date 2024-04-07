package entity

type Delivery struct {
	Id      int    `json:"-"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"-"`
	Region  string `json:"-"`
	Email   string `json:"-"`
}
