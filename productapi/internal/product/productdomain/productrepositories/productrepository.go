package productrepositories

import (
	"context"
	"strings"

	"github.com/odanaraujo/stocks-api/internal/product/productdomain/productentities"
	"gorm.io/gorm"
)

//go:generate mockery --name=ProductRepository
type ProductRepository interface {
	GetByID(_ context.Context, id string) (*productentities.Product, error)
	Search(_ context.Context, productType string) ([]*productentities.Product, error)
	Create(_ context.Context, product *productentities.Product) error
	Update(_ context.Context, productToUpdate *productentities.Product) error
	Delete(_ context.Context, id string) error
}

type productRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (p *productRepository) GetByID(_ context.Context, id string) (*productentities.Product, error) {
	product := &productentities.Product{}

	tx := p.db.Where("id = ?", id).First(product)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}

func (p *productRepository) Search(_ context.Context, productType string) ([]*productentities.Product, error) {
	limit := 10
	products := []*productentities.Product{}
	query := []string{}
	args := []interface{}{}

	query = append(query, "products.type = ?")
	args = append(args, productType)

	tx := p.db.Table("products").
		Select("products.id", "products.name", "products.type", "products.quantity")

	tx = tx.Where(strings.Join(query, " AND "), args...)

	tx = tx.Limit(limit)
	tx = tx.Scan(&products)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return products, nil

}

func (p *productRepository) Create(_ context.Context, product *productentities.Product) error {
	tx := p.db.Create(product)
	return tx.Error
}

func (p *productRepository) Update(_ context.Context, productToUpdate *productentities.Product) error {
	tx := p.db.Model(&productentities.Product{}).Where("id = ?", productToUpdate.ID).
		Updates(map[string]interface{}{
			"name":     productToUpdate.Name,
			"type":     productToUpdate.Type,
			"quantity": productToUpdate.Quantity,
		})

	return tx.Error
}

func (p *productRepository) Delete(_ context.Context, id string) error {
	tx := p.db.Where("id = ?", id).Delete(&productentities.Product{})
	return tx.Error
}
