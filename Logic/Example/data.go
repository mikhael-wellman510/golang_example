package main

import (
	"latihan-golang/Logic/Example/api"
	"latihan-golang/Logic/Example/controller"
	"latihan-golang/Logic/Example/repo"
	"latihan-golang/Logic/Example/service"
)

func main() {

	// Instance Repo
	productRepo := repo.NewProductRepository()
	userRepo := repo.NewUserRepository()
	transactionRepo := repo.NewTransactionRepo()
	transactionDetailRepo := repo.NewTransactionDetailRepository()

	// Inject repo ke Service
	productService := service.NewProductService(productRepo)
	userService := service.NewUserService(userRepo)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionDetailService := service.NewTransactionDetailService(transactionDetailRepo, userService, productService, transactionService)

	// Inject ke service ke Controller
	productController := controller.NewProductController(productService)
	userController := controller.NewUserController(userService)
	transactionDetailController := controller.NewTransactionController(transactionDetailService)

	api.HitProduct(*productController)
	api.HitUser(*userController)
	api.HitTransaction(*transactionDetailController)

	// fmt.Println("huaaa ", transactionDetailRepo.TransactionDetail)

	// cekss := transactionRepo.GetAllTrx()
	// fmt.Println("cece", cekss)
	//=====000======

}
