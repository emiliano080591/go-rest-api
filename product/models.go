package product

/*Product es la estructura de un producto de la base de datos*/
type Product struct {
	Id           int     `json:"id"`
	ProductCode  string  `json:"productCode"`
	ProductName  string  `json:"productName"`
	Description  string  `json:"description"`
	StandardCost float64 `json:"standardCost"`
	ListPrice    float64 `json:"listPrice"`
	Category     string  `json:"category"`
}

/*ProductsList es el arreglo que se va enviar de los productos*/
type ProductsList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}
