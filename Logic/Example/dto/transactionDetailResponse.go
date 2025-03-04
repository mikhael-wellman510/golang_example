package dto

type TransactionDetailResponse struct {
	ProductId    int
	ProductName  string
	Quantity     int
	PricePerUnit int
	SubTotal     int
}
