package entity

type Payment struct {
	Email      string
	Name       string
	Lastname   string
	Document   string
	Address    string
	Complement string
	City       string
	Country    string
	State      *string
	Phone      string
	CEP        string
}
