package dto

type TransactionResponse struct {
	TransactionId int
	UserId        int
	TotalPrice    int
	Details       []TransactionDetailResponse
}
