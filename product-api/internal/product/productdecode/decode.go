package productdecode

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/odanaraujo/stocks-api/product-api/internal/product/productdomain/productentities"
)

func DecodeProductFromBody(r *http.Request) (*productentities.Product, error) {
	createProduct := &productentities.Product{}

	if err := json.NewDecoder(r.Body).Decode(&createProduct); err != nil {
		return nil, err
	}

	return createProduct, nil
}

func DecodeStringIDFromURI(r *http.Request) (string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return "", errors.New("empty_id_error")
	}

	return id, nil
}

// get query params type of product
func DecodeTypeQueryString(r *http.Request) string {
	return r.URL.Query().Get("type")
}
