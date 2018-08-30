package store

import (
	"fmt"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// a successful case
func TestShouldSaveToDB(t *testing.T) {
	t.Skip()
	var stationStore StationStore
	db, mock, err := sqlmock.New()
	stationStore = NewStationStore(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	station := Station{
		City:        "Depok",
		CityCode:    "DPK",
		Code:        "DPP",
		DisplayName: "Statiun Depok",
		Island:      "Jawa",
		Name:        "Depok Statiun",
	}

	query := fmt.Sprintf(
		`
	insert into %v (
		city,
		city_code,
		code,
		create_by,
		create_time,
		display_name,
		island,
		name,
		popularity_order,
		status,
		update_by,
		update_time
	) values (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12
	)
	`,
		stationTableName,
	)

	mock.ExpectBegin()
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(
		station.City,
		station.CityCode,
		station.Code,
		station.CreateBy,
		time.Now(),
		station.DisplayName,
		station.Island,
		station.Name,
		station.PopularityOrder,
		1,
		station.UpdateBy,
		time.Now(),
	)
	mock.ExpectCommit()

	// now we execute our method
	errSaveSQL := stationStore.Add(station)

	assert.Equal(t, nil, errSaveSQL)
}
