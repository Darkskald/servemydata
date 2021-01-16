package usecases

import "servemydata/domain"

type ForGettingSpectrumById interface {
	GetSpectrumById(id string) (domain.Spectrum2d, error)
}

type GetSpectrumService struct {
	GetPort ForGettingSpectrumById
}

func NewGetSpectrumService(getter ForGettingSpectrumById) GetSpectrumService {
	return GetSpectrumService{getter}
}

func (getter GetSpectrumService) Run(id string) (domain.Spectrum2d, error) {
	return getter.GetPort.GetSpectrumById(id)
}
