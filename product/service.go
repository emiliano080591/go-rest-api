package product

type Service interface {
	GetProductById(param *getProductByIDRequest) (*Product, error)
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
