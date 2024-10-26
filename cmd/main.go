package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/odanaraujo/stock-api/internal/product/productdb"
	"github.com/odanaraujo/stock-api/internal/product/producthttp"
)

func main() {
	productdb.BuildDB()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/product/{id}", producthttp.GetProductByIDHandler)
	r.Get("/products", producthttp.SearchProductsHandler)
	r.Post("/products", producthttp.CreateProductHandler)
	r.Put("/product/{id}", producthttp.UpdateProductHandler)
	r.Delete("/product/{id}", producthttp.DeleteProductHandler)

	http.ListenAndServe(":8081", r)

}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
