package store

// A Repository connection db
type Repository struct{}

// NewRepository exported
func NewRepository() IRepository {
	return &Repository{}
}

// IDataBase interface
type IRepository interface {
	CreditAssignment(investment int32) (int32, int32, int32, error)
}
