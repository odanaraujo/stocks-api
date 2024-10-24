package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

var memoryDB map[string]*Product

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

func main() {
	BuildDB()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/product/{id}", GetProductByIDHandler)
	r.Get("/products", SearchProductsHandler)
	r.Post("/products", CreateProductHandler)
	r.Put("/product/{id}", UpdateProductHandler)
	r.Delete("/product/{id}", DeleteProductHandler)

	http.ListenAndServe(":8081", r)

}

func BuildDB() {
	startProducts := make(map[string]string)
	startProducts["camisa do Sport"] = "clothing"
	startProducts["capim dourado"] = "plant"
	startProducts["CD do dead fish"] = "music"
	startProducts["Jockey 365"] = "bar"
	startProducts["Mcbook"] = "technology"

	memoryDB = make(map[string]*Product)

	i := 0
	for product, productType := range startProducts {
		id := fmt.Sprintf("%d", i)
		memoryDB[id] = &Product{
			ID:       id,
			Name:     product,
			Type:     productType,
			Quantity: 100,
		}
		i++
	}
}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, ok := memoryDB[id]
	if !ok {
		err := errors.New("product_not_found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	WriteJSONResponse(w, product, http.StatusOK)
}

func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {

	productType := DecodeTypeQueryString(r)
	var matchedValues []*Product

	for _, value := range memoryDB {
		if value.Type == productType {
			matchedValues = append(matchedValues, value)
		}
	}

	WriteJSONResponse(w, matchedValues, http.StatusOK)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	product, err := DecodeProductFromBody(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	idString := id.String()
	product.ID = idString
	memoryDB[idString] = product
	WriteJSONResponse(w, product, http.StatusOK)

}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	currentProduct, ok := memoryDB[id]

	if !ok {
		err := errors.New("product_not_found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	newProduct, err := DecodeProductFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newProduct = SetNewProduct(currentProduct, newProduct)
	memoryDB[newProduct.ID] = newProduct
	WriteJSONResponse(w, newProduct, http.StatusOK)

}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := memoryDB[id]; !ok {
		err := errors.New("product_not_found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	delete(memoryDB, id)

	w.WriteHeader(http.StatusNoContent)

}

func WriteJSONResponse(w http.ResponseWriter, obj interface{}, status int) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "JSON")
	w.WriteHeader(status)
	w.Write(bytes)
}

// get query params type of product
func DecodeTypeQueryString(r *http.Request) string {
	return r.URL.Query().Get("type")
}

// get param with id of product
func DecodeStringIDFromURI(r *http.Request) (string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return "", errors.New("empty_id_error")
	}

	return id, nil
}

func DecodeProductFromBody(r *http.Request) (*Product, error) {
	createProduct := &Product{}

	if err := json.NewDecoder(r.Body).Decode(&createProduct); err != nil {
		return nil, err
	}

	return createProduct, nil
}

func SetNewProduct(currentProduct, newProduct *Product) *Product {
	if newProduct.Name != "" {
		currentProduct.Name = newProduct.Name
	}

	if newProduct.Quantity != 0 {
		currentProduct.Quantity = newProduct.Quantity
	}

	if newProduct.Type != "" {
		currentProduct.Type = newProduct.Type
	}

	return newProduct
}
