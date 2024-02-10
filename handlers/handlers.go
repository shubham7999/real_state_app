package handlers

import (
	"encoding/json"
	"net/http"
	"real_state/models"
	"real_state/utils"

	"github.com/gorilla/mux"
)

func UploadHouse(w http.ResponseWriter, r *http.Request) {
	var newHouse models.House
	err := json.NewDecoder(r.Body).Decode(&newHouse)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid input format")
		return
	}

	newHouse.ID = utils.GenerateID()
	models.AddHouse(newHouse)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newHouse)
}

func GetHouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	house, err := models.GetHouseByID(id)
	if err != nil {
		utils.HandleError(w, http.StatusNotFound, "House not found")
		return
	}

	json.NewEncoder(w).Encode(house)
}

func ListHouses(w http.ResponseWriter, r *http.Request) {
	houses := models.GetAllHouses()
	json.NewEncoder(w).Encode(houses)
}

func UpdateHouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	updatedHouse, err := models.UpdateHouseStatus(id, r.Body)
	if err != nil {
		utils.HandleError(w, http.StatusNotFound, "House not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedHouse)
}

func DeleteHouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := models.DeleteHouse(id)
	if err != nil {
		utils.HandleError(w, http.StatusNotFound, "House not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "deleted successfully"})

}