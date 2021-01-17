package inmem

import (
	"errors"
	"log"
	"servemydata/domain"
	"strconv"
)

type InMemoryRepo struct {
	Spectra   map[string]domain.Spectrum2d
	CurrentId int
}

//todo add tests for all use case implementations

func NewInMemoryRepo() InMemoryRepo {
	return InMemoryRepo{Spectra: make(map[string]domain.Spectrum2d), CurrentId: 0}
}

func (imr *InMemoryRepo) NewId() string {
	imr.CurrentId += 1
	return strconv.Itoa(imr.CurrentId)
}

func (imr InMemoryRepo) GetSpectrumById(id string) (domain.Spectrum2d, error) {
	if val, ok := imr.Spectra[id]; ok {
		return val, nil
	}
	return domain.Spectrum2d{}, errors.New("Unknown spectrum id.")
}

func (imr *InMemoryRepo) AddSpectrum(d domain.Spectrum2d) (domain.Spectrum2d, error) {
	newId := imr.NewId()
	log.Printf("add spectrum with id %s", newId)
	d.Id = newId
	imr.Spectra[newId] = d
	return d, nil
}

func (imr *InMemoryRepo) DeleteSpectrum(id string) (domain.Spectrum2d, error) {
	if spec, ok := imr.Spectra[id]; ok {
		delete(imr.Spectra, id)
		return spec, nil
	}
	return domain.Spectrum2d{}, errors.New("No spectrum with such id!")
}

func (imr InMemoryRepo) ListSpectra() []domain.Spectrum2d {
	temp := make([]domain.Spectrum2d, 0)

	for _, v := range imr.Spectra {
		temp = append(temp, v)
	}

	return temp
}
