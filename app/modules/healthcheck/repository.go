//go:generate mockery --name Repository --filename=repository.go

package healthcheck

// Repository is an interface to Healthcheck repository
type Repository interface {
	Ping() (bool, error)
}

type repositoryImp struct {
}

// NewRepository is a constructor to Healthcheck repository
func NewRepository() Repository {
	return &repositoryImp{}
}

func (r *repositoryImp) Ping() (bool, error) {
	return true, nil
}
