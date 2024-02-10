package models

import (
	"encoding/json"
	"errors"
	"io"
	"sync"
)

type House struct {
	ID          string `json:"id"`
	Owner       string `json:"owner"`
	Area        string `json:"area"`
	SalePrice   int    `json:"sale_price"`
	Negotiable  bool   `json:"negotiable"`
	Status      string `json:"status"`
}

var houses []House =  make([]House, 0)
var mu sync.Mutex

var (
	ErrHouseNotFound      = errors.New("house not found")
	ErrInvalidInputFormat = errors.New("invalid input format")
)

func AddHouse(house House) {
	mu.Lock()
	defer mu.Unlock()
	houses = append(houses, house)
}

func GetHouseByID(id string) (House, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, house := range houses {
		if house.ID == id {
			return house, nil
		}
	}
	return House{}, ErrHouseNotFound
}

func GetAllHouses() []House {
	mu.Lock()
	defer mu.Unlock()
	return houses
}

func UpdateHouseStatus(id string, body io.Reader) (House, error) {
	mu.Lock()
	defer mu.Unlock()
	for i, house := range houses {
		if house.ID == id {
			err := json.NewDecoder(body).Decode(&houses[i])
			if err != nil {
				return House{}, ErrInvalidInputFormat
			}
			houses[i].ID = id
			return houses[i], nil
		}
	}
	return House{}, ErrHouseNotFound
}

func ClearHouses() {
	mu.Lock()
	defer mu.Unlock()
	houses = nil
}
func DeleteHouse(id string) error {
	mu.Lock()
	defer mu.Unlock()

	index := -1
	for i, house := range houses {
		if house.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return ErrHouseNotFound
	}

	// Remove the house from the slice
	houses = append(houses[:index], houses[index+1:]...)

	return nil
}