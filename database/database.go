package shared

import "database/sql"

/*InitDB inicializa la conexion de la base de datos*/
func InitDB() *sql.DB {
	connectionString := "root:root@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)
	defer databaseConnection.Close()

	if err != nil {
		panic(err.Error())
	}

	return databaseConnection
}
