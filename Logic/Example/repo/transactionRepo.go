package repo

import "latihan-golang/Logic/Example/entity"

type TransactionRepository struct {
	Transaction []entity.Transaction
}

func NewTransactionRepo() *TransactionRepository {

	return &TransactionRepository{
		Transaction: []entity.Transaction{},
	}
}

func (t *TransactionRepository) GetAllTrx() []entity.Transaction {

	return t.Transaction
}

// Setter data
func (t *TransactionRepository) SaveTransaction(transaction entity.Transaction) entity.Transaction {

	idx := len(t.Transaction)

	data := entity.Transaction{
		Id:         idx + 1,
		UserId:     transaction.UserId,
		TotalPrice: transaction.TotalPrice,
	}

	t.Transaction = append(t.Transaction, data)

	return data

}
