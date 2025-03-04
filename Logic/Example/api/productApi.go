package api

import (
	"fmt"
	"latihan-golang/Logic/Example/controller"
	"latihan-golang/Logic/Example/dto"
)

func HitProduct(productController controller.ProductController) {

	// Find All product
	findAllProduct1 := productController.FindAllProductController()
	fmt.Println("Product Awal : ", findAllProduct1, "\n ================================")

	//---
	// todo Add Product
	addProduct := dto.ProductRequest{Nama: "Susu", Harga: 40000, Qty: 2}
	productController.AddProductController(addProduct)
	findAllProduct2 := productController.FindAllProductController()
	fmt.Println("Data Product New : ", findAllProduct2, "\n===================================")

	//-----
	// todo findProductById
	findProductById, resInfo := productController.FindByIdController(2)
	fmt.Println("Result : ", findProductById, "Dan : ", resInfo, "\n =============================")

	// Todo -> edit datproduct

	dataEdit := dto.ProductRequest{Id: 3, Nama: "Shampo", Harga: 3000, Qty: 500}
	editProduct, _ := productController.EditProductController(dataEdit)
	fmt.Println("Hasil edit : ", editProduct, "\n==========================================")

	resEdit := productController.FindAllProductController()
	fmt.Println("After edit : ", resEdit, "\n==========================")
}
