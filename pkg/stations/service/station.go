package service

import (
	"errors"
	"log"

	"github.com/zainul/sample/pkg/ownerrors"
	db "github.com/zainul/sample/pkg/stations/store"
)

// StationService ...
type StationService struct {
	Store db.StationStore
}

// NewStationService ....
func NewStationService(store db.StationStore) StationService {
	return StationService{store}
}

// Add ...
func (s *StationService) Add(station db.Station) (err error) {
	// your logic put in here

	if station.City == "" {
		return errors.New(ownerrors.CityNotBeEmpty)
	}

	errStore := s.Store.Add(station)

	if errStore != nil {
		log.Println("Error @StationService.Add ", errStore)
		err = errStore
		return
	}

	return
}
