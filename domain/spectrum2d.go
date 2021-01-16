package domain

import "time"

// Spectrum2d is the representation of a generic two-dimensional experimental dataset with an independent variable(X) and a
// dependent variable (Y). The MeasuredTime is stored as a unix timestamp. The Id is assigned by the concrete persistence
// layer whereas the Name is defined by the user.
type Spectrum2d struct {
	Id           string
	Name         string
	Measurer     string
	MeasuredTime int64 `json:"measured_time"`
	X            []float64
	Y            []float64
}

// todo test time conversion

// GetMeasuredTime converts the internal Unix time stamp representation (integer) to the  Go time.Time struct.
func (spec Spectrum2d) GetMeasuredTime() time.Time {
	return time.Unix(spec.MeasuredTime, 0)
}

// IsValid checks if the slices holding the X and Y data have the same size which should be the case for typical
// experimental data.
func (spec Spectrum2d) IsValid() bool {
	return len(spec.X) == len(spec.Y)
}
