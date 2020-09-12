package main

import (
	"fmt"
	"net/http"

	db "github.com/emiliano080591/go-rest-api/database"
	"github.com/emiliano080591/go-rest-api/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := db.InitDB()
	fmt.Println("Servidor escuchando en puerto 3000")

	var productRepository = product.NewRepository(databaseConnection)
	var productService product.Service
	productService = product.NewService(productRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHTTPHandler(productService))

	http.ListenAndServe(":3000", r)
}
