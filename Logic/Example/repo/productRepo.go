package repo

import (
	"latihan-golang/Logic/Example/entity"
)

// Kumpulan atribut nya
type ProductRepository struct {

	// Privat final product
	Product []entity.Product
}

func NewProductRepository() ProductRepository {

	// Ini membuat instance -> instance itu mengisi cetakan struct nya
	return ProductRepository{
		Product: []entity.Product{

			{Id: 1, Nama: "Indomie", Harga: 2000, Qty: 10},
			{Id: 2, Nama: "Beras", Harga: 8000, Qty: 20},
			{Id: 3, Nama: "Sabun", Harga: 1000, Qty: 40},
		},
	}

}

func (pr *ProductRepository) FindAllProductRepo() []entity.Product {

	return pr.Product
}

func (pr *ProductRepository) AddProductRepo(product entity.Product) entity.Product {

	idx := len(pr.Product) + 1

	data := entity.Product{
		Id:    idx,
		Nama:  product.Nama,
		Harga: product.Harga,
		Qty:   product.Qty,
	}

	pr.Product = append(pr.Product, data)
	return data

}

func (pr *ProductRepository) FindProductById(id int) (entity.Product, bool) {

	for _, product := range pr.Product {

		if product.Id == id {
			return product, true
		}

	}

	return entity.Product{}, false
}
