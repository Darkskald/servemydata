package usecases

import "servemydata/domain"

// ForAddingSpectrum is the port responsible for storing new spectra in a persistence layer.
type ForAddingSpectrum interface {
	AddSpectrum(d domain.Spectrum2d) (domain.Spectrum2d, error)
}

type AddSpectrumService struct {
	AddPort ForAddingSpectrum
}

func NewAddSpectrumService(adp ForAddingSpectrum) AddSpectrumService {
	return AddSpectrumService{adp}

}

func (adder AddSpectrumService) Run(d domain.Spectrum2d) (domain.Spectrum2d, error) {
	return adder.AddPort.AddSpectrum(d)
}
