package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zainul/sample/pkg/stations/store"
)

// Response ...
type Response struct {
	Data   interface{}
	Errors []Error
}

// Error ...
type Error struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// ReqAddStation ...
type ReqAddStation struct {
	City        string `json:"city"`
	CityCode    string `json:"city_code"`
	Code        string `json:"code"`
	DisplayName string `json:"display_name"`
	Island      string `json:"island"`
	Name        string `json:"name"`
}

// Add ...
func (s *Station) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req ReqAddStation
	var res Response

	errors := make([]Error, 0)

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		errors = append(errors, Error{
			ID:      "E001",
			Title:   "Mismatch parameter value",
			Message: err.Error(),
		})
		res.Errors = errors
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		log.Println(errors)
		return
	}

	errSave := stationService.Add(store.Station{
		City:        req.City,
		CityCode:    req.CityCode,
		Code:        req.Code,
		DisplayName: req.DisplayName,
		Island:      req.Island,
		Name:        req.Name,
	})

	if errSave != nil {
		errors = append(errors, Error{
			ID:      "E002",
			Title:   "Failed Data",
			Message: "Failed saving data",
		})
		res.Errors = errors
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		log.Println(errors)
		return
	}

	json.NewEncoder(w).Encode(res)
	return
}
