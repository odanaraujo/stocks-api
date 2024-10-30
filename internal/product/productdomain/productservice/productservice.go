package productservice

import (
	"context"

	"github.com/google/uuid"
	"github.com/odanaraujo/stock-api/internal/product/productdomain/productentities"
	"github.com/odanaraujo/stock-api/internal/product/productdomain/productrepositories"
)

type ProductService struct {
	productRepository productrepositories.ProductRepository
}

func New(productRepository productrepositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (p *ProductService) GetByID(ctx context.Context, id string) (*productentities.Product, error) {
	product, err := p.productRepository.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) SearchProducts(ctx context.Context, productType string) ([]*productentities.Product, error) {
	matchedValues, err := p.productRepository.Search(ctx, productType)

	if err != nil {
		return nil, err
	}

	return matchedValues, nil
}

func (p *ProductService) Create(ctx context.Context, product *productentities.Product) (*productentities.Product, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return &productentities.Product{}, err
	}

	product.ID = id.String()

	if err = p.productRepository.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) Update(ctx context.Context, productToUpdate *productentities.Product) (*productentities.Product, error) {

	product, err := p.productRepository.Update(ctx, productToUpdate)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) Delete(ctx context.Context, id string) error {

	if err := p.productRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
