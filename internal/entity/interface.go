package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	SelectById(id string) (*Order, error)
	Select() ([]Order, error)
	// GetTotal() (int, error)
}
