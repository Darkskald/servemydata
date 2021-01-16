package usecases

import "servemydata/domain"

type ForListingSpectra interface {
	ListSpectra() []domain.Spectrum2d
}

type ListSpectraService struct {
	ListPort ForListingSpectra
}

func NewListSpectraService(lister ForListingSpectra) ListSpectraService {
	return ListSpectraService{lister}
}

func (lister ListSpectraService) Run() []domain.Spectrum2d {
	return lister.ListPort.ListSpectra()
}
