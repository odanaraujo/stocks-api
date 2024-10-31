package productdb

import (
	"fmt"

	"github.com/odanaraujo/stocks-api/product-api/internal/product/productdomain/productentities"
)

var MemoryDB map[string]*productentities.Product

func BuildDB() {
	startProducts := make(map[string]string)
	startProducts["camisa do Sport"] = "clothing"
	startProducts["capim dourado"] = "plant"
	startProducts["CD do dead fish"] = "music"
	startProducts["Jockey 365"] = "bar"
	startProducts["Mcbook"] = "technology"
	startProducts["camisa do sport 2"] = "clothing"

	MemoryDB = make(map[string]*productentities.Product)

	i := 0
	for product, productType := range startProducts {
		id := fmt.Sprintf("%d", i)
		MemoryDB[id] = &productentities.Product{
			ID:       id,
			Name:     product,
			Type:     productType,
			Quantity: 100,
		}
		i++
	}
}
