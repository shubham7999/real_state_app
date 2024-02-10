package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"real_state/models"
	"testing"

	"github.com/gorilla/mux"
)
func TestUploadHouse(t *testing.T) {

	house := models.House{
		Owner:       "John Doe",
		Area:        "1500",
		SalePrice:   250000,
		Negotiable:  true,
	}

	houseJSON, _ := json.Marshal(house)

	req, _ := http.NewRequest("POST", "/api/v1/houses", bytes.NewBuffer(houseJSON))

	res := httptest.NewRecorder()

	UploadHouse(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, res.Code)
	}

	var createdHouse models.House
	json.NewDecoder(res.Body).Decode(&createdHouse)

	if createdHouse.ID == "" {
		t.Error("Expected non-empty ID")
	}

	if createdHouse.Owner != house.Owner || createdHouse.Area != house.Area ||
		createdHouse.SalePrice != house.SalePrice || createdHouse.Negotiable != house.Negotiable {
		t.Error("Expected created house to match input")
	}
}
func TestGetHouse(t *testing.T) {

	router := mux.NewRouter()

	router.HandleFunc("/houses/{id}", GetHouse).Methods("GET")

	testHouse := models.House{
		ID:          "1",
		Owner:       "John Doe",
		Area:        "Test Street",
		SalePrice:   500000,
		Negotiable:  true,
		Status:      "available",
	}

	models.AddHouse(testHouse)

	req, _ := http.NewRequest("GET", "/houses/1", nil)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var responseHouse models.House
	err := json.Unmarshal(rr.Body.Bytes(), &responseHouse)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	
}


func TestListHouses(t *testing.T) {

	models.ClearHouses()

	testHouse := models.House{
		ID:          "1",
		Owner:       "John Doe",
		Area:        "Test Street",
		SalePrice:   500000,
		Negotiable:  true,
		Status:      "available",
	}

	models.AddHouse(testHouse)

	req, _ := http.NewRequest("GET", "/houses", nil)

	rr := httptest.NewRecorder()

	ListHouses(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var responseHouses []models.House
	err := json.Unmarshal(rr.Body.Bytes(), &responseHouses)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if len(responseHouses) != 1 || responseHouses[0].ID != testHouse.ID || responseHouses[0].Owner != testHouse.Owner {
		t.Errorf("Expected list of houses to contain %+v, got %+v", testHouse, responseHouses)
	}
}

func TestUpdateHouse(t *testing.T) {

	router := mux.NewRouter()

	router.HandleFunc("/houses/{id}", UpdateHouse).Methods("PUT")

	testHouse := models.House{
		ID:          "1",
		Owner:       "John Doe",
		Area:        "Test Street",
		SalePrice:   500000,
		Negotiable:  true,
		Status:      "available",
	}

	models.AddHouse(testHouse)

	updatedData := map[string]interface{}{
		"status": "sold",
	}

	updatedDataJSON, err := json.Marshal(updatedData)
	if err != nil {
		t.Fatalf("Error encoding updated data: %v", err)
	}

	req, _ := http.NewRequest("PUT", "/houses/1", bytes.NewBuffer(updatedDataJSON))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var responseHouse models.House
	err = json.Unmarshal(rr.Body.Bytes(), &responseHouse)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if responseHouse.Status != "sold" {
		t.Errorf("Expected house status to be 'sold', got '%s'", responseHouse.Status)
	}

	if responseHouse.ID != testHouse.ID || responseHouse.Owner != testHouse.Owner {
		t.Errorf("Expected house ID and owner to remain unchanged, got %+v", responseHouse)
	}
}

func TestDeleteHouse(t *testing.T) {

	router := mux.NewRouter()


	router.HandleFunc("/houses/{id}", DeleteHouse).Methods("DELETE")


	testHouse := models.House{
		ID:          "1",
		Owner:       "John Doe",
		Area:        "Test Street",
		SalePrice:   500000,
		Negotiable:  true,
		Status:      "available",
	}

	models.AddHouse(testHouse)

	req, _ := http.NewRequest("DELETE", "/houses/1", nil)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	
}


