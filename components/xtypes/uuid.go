package xtypes

import (
	"database/sql/driver"
	"errors"

	"github.com/google/uuid"
)

type Uuid string

func (id Uuid) Value() (v driver.Value, err error) {
	if id == "" {
		return nil, nil
	}
	u, err := uuid.Parse(string(id))
	if err != nil {
		return nil, err
	}
	uuidBinary, err := u.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return uuidBinary, nil
}

func (id *Uuid) Scan(value interface{}) (err error) {
	if value == nil {
		*id = ""
		return nil
	}
	s, ok := value.([]byte)

	if !ok {
		*id = ""
		return errors.New("invalid scan source")
	}

	u, err := uuid.FromBytes(s)
	if err != nil {
		*id = ""
		return err
	}
	*id = Uuid(u.String())
	return nil
}

func (id Uuid) String() string {
	return string(id)
}

func (id Uuid) Binary() ([]byte, error) {
	u, err := uuid.Parse(string(id))
	if err != nil {
		return nil, err
	}
	uuidBinary, err := u.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return uuidBinary, nil
}

func (id Uuid) MustBinary() []byte {
	u, err := uuid.Parse(string(id))
	if err != nil {
		panic(err)
	}
	uuidBinary, err := u.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return uuidBinary
}

func NewUuid() Uuid {
	return Uuid(uuid.New().String())
}
