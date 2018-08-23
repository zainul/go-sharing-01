package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/zainul/sample/service"
	"github.com/zainul/sample/store"
)

func initDB() (map[string]*sql.DB, error) {

	var DBConnectionMap map[string]*sql.DB

	KAIMaster, err := sql.Open("postgres", "user=postgres password=ZainToped dbname=kai host=localhost port=5432 sslmode=disable")

	if err != nil {
		panic(err)
	}

	DBConnectionMap = map[string]*sql.DB{
		"db_kai.master": KAIMaster,
	}

	db := DBConnectionMap["db_kai.master"]

	errPing := db.Ping()

	if errPing != nil {
		fmt.Println("database un reachable ...")
		return DBConnectionMap, errPing
	}

	fmt.Println("database is running ....")

	return DBConnectionMap, nil
}

func main() {
	db, err := initDB()

	if err != nil {
		panic(err)
	}

	stationStore := store.NewStationStore(db["db_kai.master"])
	stationService := service.NewStationService(stationStore)

	errSave := stationService.Add(store.Station{
		City:        "Depok2",
		CityCode:    "DPK-A",
		Code:        "DPK2",
		DisplayName: "Depok2",
		Island:      "Jawa",
		Name:        "Cilangkap",
	})

	if errSave != nil {
		fmt.Println("Failed save ", errSave)
		return
	}

	fmt.Println("station Saved")

}
