package productservice

import (
	"context"

	"github.com/google/uuid"
	"github.com/odanaraujo/stocks-api/internal/product/productdomain/productentities"
	"github.com/odanaraujo/stocks-api/internal/product/productdomain/productrepositories"
)

type ProductService interface {
	GetByID(_ context.Context, id string) (*productentities.Product, error)
	Search(_ context.Context, productType string) ([]*productentities.Product, error)
	Create(_ context.Context, product *productentities.Product) (*productentities.Product, error)
	Update(_ context.Context, productToUpdate *productentities.Product) error
	Delete(_ context.Context, id string) error
}

type productService struct {
	productRepository productrepositories.ProductRepository
}

func New(productRepository productrepositories.ProductRepository) *productService {
	return &productService{productRepository: productRepository}
}

func (p *productService) GetByID(ctx context.Context, id string) (*productentities.Product, error) {
	product, err := p.productRepository.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productService) Search(ctx context.Context, productType string) ([]*productentities.Product, error) {
	matchedValues, err := p.productRepository.Search(ctx, productType)

	if err != nil {
		return nil, err
	}

	return matchedValues, nil
}

func (p *productService) Create(ctx context.Context, product *productentities.Product) (*productentities.Product, error) {
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

func (p *productService) Update(ctx context.Context, productToUpdate *productentities.Product) error {

	if err := p.productRepository.Update(ctx, productToUpdate); err != nil {
		return err
	}

	return nil
}

func (p *productService) Delete(ctx context.Context, id string) error {

	if err := p.productRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
