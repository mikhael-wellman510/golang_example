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
	// Inject repo ke Service
	productService := service.NewProductService(productRepo)
	userService := service.NewUserService(userRepo)
	// Inject ke service ke Controller
	productController := controller.NewProductController(productService)
	userController := controller.NewUserController(userService)

	api.HitProduct(productController)
	api.HitUser(userController)
	//=====000======

}
