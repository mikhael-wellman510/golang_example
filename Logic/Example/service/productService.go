package service

import (
	"latihan-golang/Logic/Example/dto"
	"latihan-golang/Logic/Example/entity"
	"latihan-golang/Logic/Example/repo"
	"time"
)

type ProductService struct {

	// Private final ProductRepository
	ProductRepository repo.ProductRepository
}

func NewProductService(productRepository repo.ProductRepository) ProductService {

	// Todo -> ini isi data ke struct , seperti this di oop
	return ProductService{ProductRepository: productRepository}
}

// Method nya
func (ps *ProductService) FindAll() []entity.Product {

	return ps.ProductRepository.FindAllProductRepo()
}

func (ps *ProductService) AddProductService(productReq dto.ProductRequest) dto.ProductResponse {

	product := entity.Product{
		Nama:  productReq.Nama,
		Harga: productReq.Harga,
		Qty:   productReq.Qty,
	}
	saveData := ps.ProductRepository.AddProductRepo(product)
	return dto.ProductResponse{
		Id:        saveData.Id,
		Nama:      saveData.Nama,
		Harga:     saveData.Harga,
		Quantity:  saveData.Qty,
		CreatedAt: time.Now(),
	}
}

func (ps *ProductService) FindProductByIdService(id int) (dto.ProductResponse, bool) {

	res, bool := ps.ProductRepository.FindProductById(id)

	if !bool {
		return dto.ProductResponse{}, false
	}

	return dto.ProductResponse{
		Id:        res.Id,
		Nama:      res.Nama,
		Harga:     res.Harga,
		Quantity:  res.Qty,
		CreatedAt: time.Now(),
	}, true
}
