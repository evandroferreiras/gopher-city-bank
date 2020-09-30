//go:generate mockery --name Repository --filename=repository.go

package account

// Repository is an interface to Account repository
type Repository interface {
	Ping() (bool, error)
}

type repositoryImp struct {
}

// NewRepository is a constructor to Account repository
func NewRepository() Repository {
	return &repositoryImp{}
}

func (r *repositoryImp) Ping() (bool, error) {
	return true, nil
}
