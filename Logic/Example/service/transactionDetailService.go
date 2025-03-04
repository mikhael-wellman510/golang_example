package service

import (
	"latihan-golang/Logic/Example/dto"
	"latihan-golang/Logic/Example/entity"
	"latihan-golang/Logic/Example/repo"
)

type TransactionDetailService struct {
	TransactionRepository *repo.TransactionDetailRepository
	TransactionService    *TransactionService
	UserService           *UserService
	ProductService        *ProductService
}

func NewTransactionDetailService(trxDetailRepository *repo.TransactionDetailRepository, userService *UserService, productService *ProductService, transactionService *TransactionService) *TransactionDetailService {

	return &TransactionDetailService{
		TransactionRepository: trxDetailRepository,
		UserService:           userService,
		ProductService:        productService,
		TransactionService:    transactionService,
	}
}

func (tds *TransactionDetailService) AddTransactionDetail(transactionReq dto.TransactionRequest) dto.TransactionResponse {
	/*
		dapatkan User Id
	*/

	user, _ := tds.UserService.FindByIdService(transactionReq.UserId)

	totalHarga := 0

	for _, val := range transactionReq.TransactionDetailList {
		data, _ := tds.ProductService.FindProductByIdService(val.ProductId)
		// Maping ke transaction
		totalHarga = totalHarga + (data.Harga * val.Qty)

	}
	//Ssave to Db Transaction
	transactions := entity.Transaction{
		UserId:     user.Id,
		TotalPrice: totalHarga,
	}

	resTrx := tds.TransactionService.AddTransaction(transactions)
	listTransaction := []dto.TransactionDetailResponse{}

	subTotal := 0
	for _, vals := range transactionReq.TransactionDetailList {
		data, _ := tds.ProductService.FindProductByIdService(vals.ProductId)
		saveTransactionDetail := entity.TransactionDetail{
			TransactionId: resTrx.Id,
			ProductId:     vals.ProductId,
			Quantity:      vals.Qty,
			PricePerUnit:  data.Harga,
			SubTotal:      data.Harga * vals.Qty,
		}
		subTotal = subTotal + saveTransactionDetail.SubTotal
		saveData := tds.TransactionRepository.SaveTransactionDetail(saveTransactionDetail)
		mappingTransactionDetailResponse := dto.TransactionDetailResponse{
			ProductId:    saveData.ProductId,
			ProductName:  data.Nama,
			Quantity:     saveData.Quantity,
			PricePerUnit: saveData.PricePerUnit,
			SubTotal:     saveData.SubTotal,
		}

		listTransaction = append(listTransaction, mappingTransactionDetailResponse)

	}

	return dto.TransactionResponse{
		TransactionId: resTrx.Id,
		UserId:        user.Id,
		TotalPrice:    subTotal,
		Details:       listTransaction,
	}
}
