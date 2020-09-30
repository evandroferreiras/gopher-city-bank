//go:generate mockery --name Service --filename=service.go

package healthcheck

// Service is an interface to Healthcheck service
type Service interface {
	IsWorking() bool
}

type serviceImp struct {
	repository Repository
}

// NewService is a constructor to Healthcheck service
func NewService() Service {
	return &serviceImp{
		repository: NewRepository(),
	}
}

// IsWorking returns true or false, depending on what database ping return.
func (s *serviceImp) IsWorking() bool {
	isOk, err := s.repository.Ping()
	if err != nil {
		return false
	}
	return isOk
}
