package producthttp

import (
	"net/http"

	"github.com/odanaraujo/stock-api/internal/encode"
	"github.com/odanaraujo/stock-api/internal/product/productdecode"
	"github.com/odanaraujo/stock-api/internal/product/productdomain/productservice"
)

var productService = productservice.New()

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := productService.GetByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJSONResponse(w, product, http.StatusOK)
}

func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	productType := productdecode.DecodeTypeQueryString(r)

	matchedValues, err := productService.SearchProducts(ctx, productType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJSONResponse(w, matchedValues, http.StatusOK)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	product, err := productdecode.DecodeProductFromBody(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err = productService.Create(ctx, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	encode.WriteJSONResponse(w, product, http.StatusOK)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	productToUpdate, err := productdecode.DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productToUpdate, err = productService.Update(ctx, productToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJSONResponse(w, productToUpdate, http.StatusOK)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := productdecode.DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = productService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encode.WriteJSONResponse(w, nil, http.StatusNoContent)
}
