package entity

type Company struct {
	// Define the fields returned by /basic endpoint
	Name    string `json:"name"`
	Address string `json:"address"`
	// Add other fields as needed
}
