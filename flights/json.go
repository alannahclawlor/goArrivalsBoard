package flights

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Struct for flight status
type FStatus struct {
	Code      string `json:"code"`
	Arrived   string `json:"arrived,omitempty"`
	Expected  string `json:"expected_at,omitempty"`
	Cancelled bool   `json:"cancelled"`
}

// Struct to represent a single flight data entry
type JFlight struct {
	Code   string    `json:"code"`
	From   string    `json:"from"`
	Time   time.Time `json:"scheduled_arrival"`
	Status FStatus   `json:"status"`
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

	// Decodes the JSON file
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return data.Flights, nil
}

// Function to print flight data
func PrintFlights(Flights []JFlight) {
	fmt.Println("Time From Code Status")
	for _, flight := range Flights {
		status := GetStatus(flight)
		formattedTime := flight.Time.Format("15:04")
		fmt.Printf("%s %s %s %s\n", formattedTime, flight.From, flight.Code, status)
	}
}

func GetStatus(f JFlight) string {
	if f.Status.Cancelled {
		return "Cancelled"
	} else if f.Status.Arrived != "" {
		return string("Landed " + f.Status.Arrived[11:16])
	} else {
		return string("Expected " + f.Status.Expected[11:16])
	}
}
