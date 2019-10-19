package usecases

// ServoInteractor struct
type ServoInteractor struct {
	ServoRepository ServoRepository
}

// ServoRepository interface
type ServoRepository interface {
	Find() (string, error)
}

// Find function
func (interactor *ServoInteractor) Find() (string, error) {
	return interactor.ServoRepository.Find()
}
