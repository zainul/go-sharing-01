package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/zainul/sample/internal"
	st "github.com/zainul/sample/pkg/stations/handler"
)

func main() {

	db := internal.GetDB()

	st.NewStationHandler(db)
	stationHandler := st.Station{}

	r := mux.NewRouter()
	r.HandleFunc("/stations", stationHandler.Add).Methods("POST")

	http.Handle("/", r)
	log.Println("Server up :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
