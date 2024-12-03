package productentities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetNewProduct_WithSuccess(t *testing.T) {
	product := Product{
		ID:       "1",
		Name:     "teste1",
		Type:     "food",
		Quantity: 10,
	}

	newProduct := Product{
		ID:       "1",
		Name:     "teste2",
		Type:     "food",
		Quantity: 10,
	}

	productCreate := SetNewProduct(&product, &newProduct)

	assert.Equal(t, productCreate.ID, product.ID)
	assert.Equal(t, productCreate.Name, product.Name)
}
