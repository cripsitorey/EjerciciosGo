package medicine

import (
	"errors"
	"time"
)

const DateLayout = "2006-01-02"
const DisplayLayout = "02/01/2006"

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int
}

func NewMedicine(name string, manufacturer string, manufactureDate time.Time, shelfLife int) (*Medicine, error) {
	m := &Medicine{}

	if err := m.SetName(name); err != nil {
		return nil, err
	}
	if err := m.SetManufacturer(manufacturer); err != nil {
		return nil, err
	}
	if err := m.SetManufactureDate(manufactureDate); err != nil {
		return nil, err
	}
	if err := m.SetShelfLife(shelfLife); err != nil {
		return nil, err
	}

	return m, nil
}

func (m Medicine) GetName() string {
	return m.name
}

func (m Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m Medicine) GetDate() time.Time {
	return m.manufactureDate
}

func (m Medicine) GetShelf() int {
	return m.shelfLife
}

func (m *Medicine) SetName(name string) error {
	if name == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	m.name = name
	return nil
}

func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return errors.New("el fabricante no puede estar vacío")
	}
	m.manufacturer = manufacturer
	return nil
}

func (m *Medicine) SetManufactureDate(date time.Time) error {
	if date.IsZero() {
		return errors.New("la fecha de fabricación es obligatoria")
	}
	if date.After(time.Now()) {
		return errors.New("la fecha de fabricación no puede ser futura")
	}
	m.manufactureDate = date
	return nil
}

func (m *Medicine) SetShelfLife(months int) error {
	if months <= 0 {
		return errors.New("la vida útil debe ser un número positivo de meses")
	}
	m.shelfLife = months
	return nil
}

func (m *Medicine) ExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

func (m *Medicine) IsExpired() bool {
	return time.Now().After(m.ExpirationDate())
}
