package service

import (
	"errors"
	"log"

	db "github.com/zainul/sample/store"
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
		return errors.New("City not be empty")
	}

	errStore := s.Store.Add(station)

	if errStore != nil {
		log.Println("Error @StationService.Add ", errStore)
		err = errStore
		return
	}

	return
}
