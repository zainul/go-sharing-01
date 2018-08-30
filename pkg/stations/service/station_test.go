package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zainul/sample/pkg/ownerrors"
	"github.com/zainul/sample/pkg/stations/store"
)

func TestStationCityEmpty(t *testing.T) {
	stationMock := new(store.StationMock)
	ss := NewStationService(stationMock)

	err := ss.Add(store.Station{})

	assert.NotEqual(t, nil, err)
	assert.Equal(t, ownerrors.CityNotBeEmpty, err.Error())
}

func TestStationValid(t *testing.T) {
	stationMock := new(store.StationMock)
	ss := NewStationService(stationMock)

	stationData := store.Station{
		City:        "Depok2",
		CityCode:    "DPK-A",
		Code:        "DPK2",
		DisplayName: "Depok2",
		Island:      "Jawa",
		Name:        "Cilangkap",
	}

	stationMock.On("Add", stationData).Return(nil)
	err := ss.Add(stationData)

	assert.Equal(t, nil, err)
}

func TestStationInValid(t *testing.T) {
	stationMock := new(store.StationMock)
	ss := NewStationService(stationMock)

	stationData := store.Station{
		City:        "Error",
		CityCode:    "DPK-A",
		Code:        "DPK2",
		DisplayName: "Depok2",
		Island:      "Jawa",
		Name:        "Cilangkap",
	}

	stationMock.On("Add", stationData).Return(errors.New("got some error"))
	err := ss.Add(stationData)

	assert.NotEqual(t, nil, err)
}
