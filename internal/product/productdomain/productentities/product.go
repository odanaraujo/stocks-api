package productentities

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

func SetNewProduct(currentProduct, newProduct *Product) *Product {
	if newProduct.Name != "" {
		currentProduct.Name = newProduct.Name
	}

	if newProduct.Quantity != 0 {
		currentProduct.Quantity = newProduct.Quantity
	}

	if newProduct.Type != "" {
		currentProduct.Type = newProduct.Type
	}

	return newProduct
}
