package product

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error)
	GetPoducts(params *getProductsRequest) (*ProductsList, error)
}

/*Service es ña interface que va descodificar de la url*/
type service struct {
	repo Repository
}

/*NewService hace referencia al repository de repository.go*/
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

/*GetProductById hace la logica de negocio*/
func (s *service) GetProductById(param *getProductByIDRequest) (*Product, error) {
	//Business Logic
	return s.repo.GetProductById(param.ProductID)
}

func (s *service) GetPoducts(params *getProductsRequest) (*ProductsList, error) {
	products, err := s.repo.GetProducts(params)
	if err != nil {
		panic(err)
	}
	totalProducts, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}
	return &ProductsList{Data: products, TotalRecords: totalProducts}, nil
}
