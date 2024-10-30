package productrepositories

import (
	"context"
	"errors"

	"github.com/odanaraujo/stock-api/internal/product/productdb"
	"github.com/odanaraujo/stock-api/internal/product/productdomain/productentities"
)

type ProductRepository struct{}

func New() *ProductRepository {
	return &ProductRepository{}
}

func (p *ProductRepository) GetByID(_ context.Context, id string) (*productentities.Product, error) {
	product, ok := productdb.MemoryDB[id]
	if !ok {
		return nil, errors.New("product_not_found")
	}

	return product, nil
}

func (p *ProductRepository) Search(_ context.Context, productType string) ([]*productentities.Product, error) {
	var matchedValues []*productentities.Product
	for _, value := range productdb.MemoryDB {
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}

	return matchedValues, nil
}

func (p *ProductRepository) Create(_ context.Context, product *productentities.Product) error {
	productdb.MemoryDB[product.ID] = product
	return nil
}

func (p *ProductRepository) Update(_ context.Context, productToUpdate *productentities.Product) (*productentities.Product, error) {
	product, ok := productdb.MemoryDB[productToUpdate.ID]
	if !ok {
		return nil, errors.New("product_not_found")
	}

	product = productentities.SetNewProduct(product, productToUpdate)
	productdb.MemoryDB[productToUpdate.ID] = product
	return product, nil
}

func (p *ProductRepository) Delete(_ context.Context, id string) error {
	if _, ok := productdb.MemoryDB[id]; !ok {
		return errors.New("product_not_found")
	}

	delete(productdb.MemoryDB, id)
	return nil
}
