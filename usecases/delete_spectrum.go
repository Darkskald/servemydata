package usecases

import "servemydata/domain"

type ForDeletingSpectrum interface {
	DeleteSpectrum(id string) (domain.Spectrum2d, error)
}

type DeleteSpectrumService struct {
	DeletePort ForDeletingSpectrum
}

func NewDeleteSpectrumService(deleter ForDeletingSpectrum) DeleteSpectrumService {
	return DeleteSpectrumService{deleter}
}

func (deleter DeleteSpectrumService) Run(id string) (domain.Spectrum2d, error) {
	return deleter.DeletePort.DeleteSpectrum(id)
}
