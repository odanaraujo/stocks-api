package productservice

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/odanaraujo/stock-api/internal/product/productdb"
	"github.com/odanaraujo/stock-api/internal/product/productdomain/productentities"
)

type ProductService struct {
}

func New() *ProductService {
	return &ProductService{}
}

func (p *ProductService) GetByID(_ context.Context, id string) (*productentities.Product, error) {
	product, ok := productdb.MemoryDB[id]
	if !ok {
		return nil, errors.New("product_not_found")
	}

	return product, nil
}

func (p *ProductService) SearchProducts(_ context.Context, productType string) ([]*productentities.Product, error) {
	var matchedValues []*productentities.Product
	for _, value := range productdb.MemoryDB {
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}

	return matchedValues, nil
}

func (p *ProductService) Create(_ context.Context, product *productentities.Product) (*productentities.Product, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return &productentities.Product{}, err
	}

	idString := id.String()
	product.ID = id.String()

	productdb.MemoryDB[idString] = product
	return product, nil
}

func (p *ProductService) Update(_ context.Context, productToUpdate *productentities.Product) (*productentities.Product, error) {
	product, ok := productdb.MemoryDB[productToUpdate.ID]
	if !ok {
		return &productentities.Product{}, errors.New("product_not_found")
	}

	product = productentities.SetNewProduct(product, productToUpdate)
	productdb.MemoryDB[productToUpdate.ID] = product
	return product, nil
}

func (p *ProductService) Delete(_ context.Context, id string) error {

	if _, ok := productdb.MemoryDB[id]; !ok {
		return errors.New("product_not_found")
	}

	delete(productdb.MemoryDB, id)
	return nil
}
