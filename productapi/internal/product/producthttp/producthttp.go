package producthttp

import (
	"net/http"

	"github.com/odanaraujo/stocks-api/internal/encode"
	"github.com/odanaraujo/stocks-api/internal/product/productdecode"
	"github.com/odanaraujo/stocks-api/internal/product/productdomain/productservice"
)

type ProducHttp interface{}

type producHttp struct {
	productService productservice.ProductService
}

func New(productService productservice.ProductService) *producHttp {
	return &producHttp{productService: productService}
}

func (p *producHttp) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
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

func (p *producHttp) SearchProductsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	productType := productdecode.DecodeTypeQueryString(r)

	matchedValues, err := p.productService.Search(ctx, productType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJSONResponse(w, matchedValues, http.StatusOK)
}

func (p *producHttp) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
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

func (p *producHttp) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
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

func (p *producHttp) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
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
