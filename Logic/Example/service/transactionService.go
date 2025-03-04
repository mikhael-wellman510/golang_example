package service

import (
	"latihan-golang/Logic/Example/entity"
	"latihan-golang/Logic/Example/repo"
)

type TransactionService struct {
	TransactionRepository *repo.TransactionRepository
}

func NewTransactionService(transacation *repo.TransactionRepository) *TransactionService {

	return &TransactionService{
		TransactionRepository: transacation,
	}
}

func (ts *TransactionService) AddTransaction(trx entity.Transaction) entity.Transaction {

	return ts.TransactionRepository.SaveTransaction(trx)

}
