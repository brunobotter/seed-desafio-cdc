package entity

type Payment struct {
	CustomerId int64
	Coupon     string
	Total      float64
	Itens      []Itens
}

type Itens struct {
	BookId int64
	Amount int64
	Price  float64
}
