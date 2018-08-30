package handler

import (
	"database/sql"

	"github.com/zainul/sample/pkg/stations/service"
	"github.com/zainul/sample/pkg/stations/store"
)

const kaiMaster = "db_kai.master"

var (
	stationStore   store.StationStore
	stationService service.StationService
)

// Station handler
type Station struct{}

// NewStationHandler ...
func NewStationHandler(db map[string]*sql.DB) {
	stationStore = store.NewStationStore(db[kaiMaster])
	stationService = service.NewStationService(stationStore)
}
