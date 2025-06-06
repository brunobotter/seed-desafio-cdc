package entity

type Payment struct {
	CustomerId int64
	Total      float64
	Itens      []Itens
}

type Itens struct {
	BookId int64
	Amount int64
	Price  float64
}
