package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/emiliano080591/go-rest-api/database"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

var databaseConnection *sql.DB

/*Product prototipo de un producto*/
type Product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"product_code"`
	Description string `json:"description"`
}

func main() {
	databaseConnection = db.InitDB()
	defer databaseConnection.Close()
	fmt.Println("Servidor escuchando en localhost:3000")
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/products", AllProductos)
	r.Post("/products", CreateProducto)
	r.Put("/products/{id}", UpdateProducto)
	r.Delete("/products/{id}", DeleteProducto)

	http.ListenAndServe(":3000", r)
	fmt.Println("Servidor escuchando en localhost:3000")
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

/*AllProductos obtiene todos los productos*/
func AllProductos(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id,product_code,COALESCE(description,'') FROM products`

	results, err := databaseConnection.Query(sql)
	catch(err)

	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.ProductCode, &product.Description)

		catch(err)
		products = append(products, product)
	}

	respondwithJSON(w, http.StatusOK, products)
}

/*CreateProducto crea un producto nuevo en la base de datos*/
func CreateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := databaseConnection.Prepare("INSERT products SET product_code=?, description=?")
	catch(err)

	_, er := query.Exec(producto.ProductCode, producto.Description)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "succesfully created"})
}

/*UpdateProducto actualiza un producto nuevo en la base de datos*/
func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := databaseConnection.Prepare("UPDATE products SET product_code=?, description=? WHERE id=?")
	catch(err)

	_, er := query.Exec(producto.ProductCode, producto.Description, id)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "succesfully update"})
}

/*DeleteProducto elimina un producto nuevo en la base de datos*/
func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := databaseConnection.Prepare("DELETE FROM products WHERE id=?")
	catch(err)

	_, er := query.Exec(id)
	catch(er)
	defer query.Close()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "succesfully delete"})
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
