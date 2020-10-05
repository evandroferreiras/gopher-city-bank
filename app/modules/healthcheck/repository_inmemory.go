package healthcheck

type repositoryImp struct {
}

// NewInMemoryDBRepository is a constructor to inmemory Healthcheck repository
func NewInMemoryDBRepository() Repository {
	return &repositoryImp{}
}

func (r *repositoryImp) Ping() (bool, error) {
	return true, nil
}
