package main

import (
	"fmt"

	db "github.com/emiliano080591/go-rest-api/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := db.InitDB()
	fmt.Println(databaseConnection)
	//r := chi.NewRouter()
	//r.Use(middleware.Logger)
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("welcome"))
	//})
	//http.ListenAndServe(":3000", r)
}
