package medicine

import (
	"errors"
	"time"
)

type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}

func NewTablet(name string, manufacturer string, manufactureDate time.Time, shelfLife int,
	dosePerTablet float64, isPrescription bool) (*Tablet, error) {

	base, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return nil, err
	}

	t := &Tablet{
		Medicine:       *base,
		isPrescription: isPrescription,
	}

	if err := t.SetDosePerTablet(dosePerTablet); err != nil {
		return nil, err
	}

	return t, nil
}

func (t Tablet) DosePerTablet() float64 {
	return t.dosePerTablet
}

func (t Tablet) IsPrescription() bool {
	return t.isPrescription
}

func (t *Tablet) SetDosePerTablet(dose float64) error {
	if dose <= 0 {
		return errors.New("la dosis por tableta debe ser mayor a 0 mg")
	}
	t.dosePerTablet = dose
	return nil
}

func (t *Tablet) SetIsPrescription(value bool) {
	t.isPrescription = value
}
