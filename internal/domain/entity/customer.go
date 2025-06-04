package entity

type Customer struct {
	Email      string
	Name       string
	Lastname   string
	Document   string
	Address    string
	Complement string
	City       string
	Country    int64
	State      *int64
	Phone      string
	CEP        string
}
