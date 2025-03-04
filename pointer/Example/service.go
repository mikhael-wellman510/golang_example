package example

type Product struct {
	Nama string
	Qty  int
}

type ProductList struct {
	DataProduct []Product
}

func NewData() ProductList {

	return ProductList{}
}

func (pl *ProductList) AddProduct(data Product) {
	pl.DataProduct = append(pl.DataProduct, data)

}

func (pl *ProductList) FindAll() []Product {

	return pl.DataProduct
}
