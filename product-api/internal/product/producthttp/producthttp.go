package producthttp

import (
	"net/http"

	"github.com/odanaraujo/stocks-api/product-api/internal/encode"
	"github.com/odanaraujo/stocks-api/product-api/internal/product/productdecode"
	"github.com/odanaraujo/stocks-api/product-api/internal/product/productdomain/productservice"
)

type ProducHttp struct {
	productService *productservice.ProductService
}

func New(productService *productservice.ProductService) *ProducHttp {
	return &ProducHttp{productService: productService}
}

func (p *ProducHttp) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := p.productService.GetByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJSONResponse(w, product, http.StatusOK)
}

func (p *ProducHttp) SearchProductsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	productType := productdecode.DecodeTypeQueryString(r)

	matchedValues, err := p.productService.SearchProducts(ctx, productType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJSONResponse(w, matchedValues, http.StatusOK)
}

func (p *ProducHttp) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	product, err := productdecode.DecodeProductFromBody(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err = p.productService.Create(ctx, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encode.WriteJSONResponse(w, product, http.StatusOK)
}

func (p *ProducHttp) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	productToUpdate, err := productdecode.DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = p.productService.Update(ctx, productToUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJSONResponse(w, productToUpdate, http.StatusOK)
}

func (p *ProducHttp) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = p.productService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJSONResponse(w, nil, http.StatusNoContent)
}
