package entity

type TransactionDetail struct {
	Id            int
	TransactionId int
	ProductId     int
	Quantity      int
	PricePerUnit  int
	SubTotal      int
}
