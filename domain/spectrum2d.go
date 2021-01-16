package domain

// Spectrum2d is the representation of a generic two-dimensional experimental dataset with an independent variable(X) and a
// dependent variable (Y). The MeasuredTime is stored as a unix timestamp. The Id is assigned by the concrete persistence
// layer whereas the Name is defined by the user.
type Spectrum2d struct {
	Id           string
	Name         string
	Measurer     string
	MeasuredTime int
	X            []float64
	Y            []float64
}

// todo ensuure MeasuredTime is documented as unix timestamp
// check if X and Y have the same dimension
