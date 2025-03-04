package controller

import (
	"latihan-golang/Logic/Example/dto"
	"latihan-golang/Logic/Example/service"
	"time"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (pc *ProductController) FindAllProductController() []dto.ProductResponse {
	productDto := []dto.ProductResponse{}

	for _, val := range pc.ProductService.FindAll() {
		productDto = append(productDto, dto.ProductResponse{
			Id:        val.Id,
			Nama:      val.Nama,
			Harga:     val.Harga,
			Quantity:  val.Qty,
			CreatedAt: time.Now(),
		})
	}
	return productDto
}

func (pc *ProductController) AddProductController(productReq dto.ProductRequest) dto.ProductResponse {

	res := pc.ProductService.AddProductService(productReq)

	return res
}

func (pc *ProductController) FindByIdController(id int) (dto.ProductResponse, string) {
	res, bool := pc.ProductService.FindProductByIdService(id)

	if !bool {
		return dto.ProductResponse{}, "Data not found"
	}
	return res, "Succes Find Data"
}

func (pc *ProductController) EditProductController(productReq dto.ProductRequest) (dto.ProductResponse, bool) {
	res, bool := pc.ProductService.EditProductService(productReq)

	if bool {
		return res, true
	}

	return res, false

}
