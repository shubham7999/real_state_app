package utils

import "fmt"

var (
	globalId = 0
)
func InitDatabase() {
	// Initialize your database connection here
	fmt.Println("Database connected successfully")
}

func GenerateID() string {
	// Implement your logic to generate a unique ID
	globalId++
	return fmt.Sprintf("%v", globalId)
}
