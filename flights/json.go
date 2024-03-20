package flights

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Struct to represent a single flight data entry
type JFlight struct {
	Code string    `json:"code"`
	From string    `json:"from"`
	Time time.Time `json:"scheduled_arrival"`
}

func ReadJSONFile(filename string) ([]JFlight, error) {
	var data struct {
		Flights []JFlight `json:"flights"`
	}

	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decodes the json file
	decoder := json.NewDecoder(file)

	// adds the json data to the private data struct to access the flight key for the nested flight objects
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return data.Flights, nil
}

// Function to print flight data
func PrintFlights(Flights []JFlight) {
	fmt.Println("Time From Code")
	for _, flight := range Flights {
		formattedTime := flight.Time.Format("15:04")
		fmt.Printf("%s %s %s\n", formattedTime, flight.From, flight.Code)
	}
}
