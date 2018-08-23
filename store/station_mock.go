package store

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

// StationMock ...
type StationMock struct {
	mock.Mock
}

// Add ...
func (s *StationMock) Add(st Station) error {

	s.Called(st)
	if st.City == "Error" {
		return errors.New("got some error")
	}
	return nil
}
