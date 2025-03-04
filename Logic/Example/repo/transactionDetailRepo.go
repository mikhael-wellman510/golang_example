package repo

import "latihan-golang/Logic/Example/entity"

type TransactionDetailRepository struct {
	TransactionDetail []entity.TransactionDetail
}

func NewTransactionDetailRepository() *TransactionDetailRepository {
	return &TransactionDetailRepository{}
}

func (td *TransactionDetailRepository) SaveTransactionDetail(data entity.TransactionDetail) entity.TransactionDetail {

	idx := len(td.TransactionDetail) + 1

	datas := entity.TransactionDetail{
		Id:            idx,
		TransactionId: data.TransactionId,
		ProductId:     data.ProductId,
		Quantity:      data.Quantity,
		PricePerUnit:  data.PricePerUnit,
		SubTotal:      data.SubTotal,
	}

	td.TransactionDetail = append(td.TransactionDetail, datas)

	return datas
}
