package medicine

import (
	"time"
)

type Syrup struct {
	Medicine
	volume float64
	flavor string
}

func NewSyrup(name string, manufacturer string, manufactureDate time.Time, shelfLife int,
	dosePerSyrup float64, isPrescription bool) (*Syrup, error) {

}

func (s Syrup) Volume() float64 {
	return s.volume
}

func (s Syrup) Flavor() string {
	return s.flavor
}
