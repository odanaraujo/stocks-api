package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/odanaraujo/stocks-api/internal/config"
	"github.com/odanaraujo/stocks-api/internal/mysql"
	"github.com/odanaraujo/stocks-api/internal/product/productdb"
	"github.com/odanaraujo/stocks-api/internal/product/productdomain/productrepositories"
	"github.com/odanaraujo/stocks-api/internal/product/productdomain/productservice"
	"github.com/odanaraujo/stocks-api/internal/product/producthttp"
)

func main() {

	cfg, err := config.Load(os.Args)

	if err != nil {
		panic(err)
	}

	db, err := mysql.Start(cfg.Mysql.Url, cfg.Mysql.Db, cfg.Mysql.User, cfg.Mysql.Password)
	if err != nil {
		panic(err)
	}

	productRepository := productrepositories.New(db)
	productService := productservice.New(productRepository)
	productHttp := producthttp.New(productService)

	productdb.BuildDB()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/product/{id}", productHttp.GetProductByIDHandler)
	r.Get("/products", productHttp.SearchProductsHandler)
	r.Post("/products", productHttp.CreateProductHandler)
	r.Put("/product/{id}", productHttp.UpdateProductHandler)
	r.Delete("/product/{id}", productHttp.DeleteProductHandler)

	http.ListenAndServe(":8081", r)

}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
