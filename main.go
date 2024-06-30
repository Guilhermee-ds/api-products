package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

func main() {
	BuilDb()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products/{id}", GetProductsByIDHandler)
	r.Get("/products", SearchProductsHandler)
	r.Post("/products", CreateProductHandler)
	r.Put("/products/{id}", UpdateProductHandler)
	r.Delete("/products/{id}", DeleteProductHandler)

	http.ListenAndServe(":8081", r)
}

var memoryDb map[string]*Product

func BuilDb() {
	startProducts := make(map[string]string)
	startProducts["Camisa do Corinthians"] = "clothing"
	startProducts["Capim Dourado"] = "plant"
	startProducts["CD do Atitude 67"] = "music"
	startProducts["Flash 165"] = "boat"
	startProducts["Bandana Dazaranha"] = "clothing"
	startProducts["Motul 5w40"] = "oil"

	memoryDb = make(map[string]*Product)

	i := 0

	for product, ProductType := range startProducts {
		id := fmt.Sprintf("%d", i)
		memoryDb[id] = &Product{
			ID:       id,
			Name:     product,
			Type:     ProductType,
			Quantity: 100,
		}
		i++
	}
}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "admin")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type HandlerFunc func(ResponseWriter, *Request)

func (mx *Mux) Get(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mGET, pattern, handlerFn)
}

func GetProductByHandler(w http.ResponseWriter, r *http.Request) {
	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
