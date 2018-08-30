package store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

var total = 0

const stationTableName = "kai_stations"

// Station station schema
type Station struct {
	City            string     `json:"city"`
	CityCode        string     `json:"city_code"`
	Code            string     `json:"code"`
	CreateBy        int        `json:"create_by"`
	CreateTime      string     `json:"create_time"`
	DeleteBy        *int       `json:"delete_by"`
	DeleteTime      *time.Time `json:"delete_time"`
	DisplayName     string     `json:"display_name"`
	ID              int        `json:"id"`
	Island          string     `json:"island"`
	Name            string     `json:"name"`
	PopularityOrder int        `json:"popularity_order"`
	Status          int        `json:"status"`
	UpdateBy        *int       `json:"update_by"`
	UpdateTime      *time.Time `json:"update_time"`
}

// StationStore ...
type StationStore interface {
	Add(station Station) error
}

type storeStation struct {
	db *sql.DB
}

// NewStationStore ...
func NewStationStore(db *sql.DB) StationStore {
	return &storeStation{db}
}

func (s *storeStation) Add(station Station) (err error) {

	tx, errTx := s.db.Begin()

	if errTx != nil {
		err = errTx

		return
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

	stmt, errPrepare := s.db.Prepare(query)

	if errPrepare != nil {
		log.Println("Error Store @Station.Add -> prepare", errPrepare)
		err = errors.New("failed to prepare the query")
		errRollback := tx.Rollback()
		logError("Error rollback prepare query", errRollback)
		return
	}
	_, errStmt := stmt.Exec(
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

	defer func() {
		err = stmt.Close()
	}()

	if errStmt != nil {
		log.Println("Error Store @Station.Add -> statement", errStmt)
		err = errStmt
		total = total + 1
		errRollback := tx.Rollback()
		logError("Error rollback stmt", errRollback)
		return
	}

	// time.Sleep(100 * time.Millisecond)

	err = tx.Commit()

	if err != nil {
		errRollback := tx.Rollback()
		logError("Failed to commit ", errRollback)
	}

	log.Println("Saved station success ... total error ", total)

	return
}

func logError(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
	}
}
