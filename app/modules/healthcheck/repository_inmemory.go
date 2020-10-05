package healthcheck

type repositoryInMemory struct {
}

// NewInMemoryDBRepository is a constructor to inmemory Healthcheck repository
func NewInMemoryDBRepository() Repository {
	return &repositoryInMemory{}
}

func (r *repositoryInMemory) Ping() (bool, error) {
	return true, nil
}
