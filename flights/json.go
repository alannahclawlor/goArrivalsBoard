package flights

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Struct for flight status
type FStatus struct {
	Code      string    `json:"code"`
	Arrived   time.Time `json:"arrived"`
	Expected  time.Time `json:"expected_at"`
	Cancelled bool      `json:"cancelled"`
}

type JStatus struct {
	Code      string `json:"code"`
	Arrived   string `json:"arrived"`
	Expected  string `json:"expected_at"`
	Cancelled bool   `json:"cancelled"`
}

// Struct to represent a single flight data entry
type JFlight struct {
	Code   string    `json:"code"`
	From   string    `json:"from"`
	Time   time.Time `json:"scheduled_arrival"`
	Status FStatus
}

// Function to convert JStatus to FStatus
func (j JStatus) JStatusConvert() FStatus {
	fStatus := FStatus{}

	// Convert 'Arrived' string to time if not empty
	if j.Arrived != "" {
		arrivedTime, err := time.Parse("2006-01-02T15:04:05Z", j.Arrived)
		if err != nil {
			fmt.Printf("Error parsing arrived time string: %v\n", err)
		} else {
			fStatus.Arrived = arrivedTime
		}
	} else {
		fStatus.Arrived = time.Time{} // Set to zero time if empty
	}

	// Convert 'Expected' string to time if not empty
	if j.Expected != "" {
		expectedTime, err := time.Parse("2006-01-02T15:04:05Z", j.Expected)
		if err != nil {
			fmt.Printf("Error parsing expected time string: %v\n", err)
		} else {
			fStatus.Expected = expectedTime
		}
	} else {
		fStatus.Expected = time.Time{} // Set to zero time if empty
	}

	fStatus.Cancelled = j.Cancelled

	return fStatus
}

func ReadJSONFile(filename string) ([]JFlight, error) {
	var data struct {
		Flights []struct {
			Code   string `json:"code"`
			From   string `json:"from"`
			Time   string `json:"scheduled_arrival"`
			Status JStatus
		} `json:"flights"`
	}

	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open Json error: ", err)
		return nil, err
	}
	defer file.Close()

	// Decodes the json file
	decoder := json.NewDecoder(file)

	// Decode JSON data into 'data' struct
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Decode Json error: ", err)
		return nil, err
	}

	// Create a slice to hold JFlight instances
	var flights []JFlight

	// Iterate over the flights data and populate JFlight instances
	for _, flightData := range data.Flights {
		// Convert the time string to a time.Time object
		scheduledArrival, err := time.Parse(time.RFC3339, flightData.Time)
		if err != nil {
			fmt.Printf("Error parsing time for flight %s: %v\n", flightData.Code, err)
			continue // Skip this flight if time parsing fails
		}

		// Convert the status data to FStatus
		status := flightData.Status.JStatusConvert()

		// Create a new JFlight instance and append it to the flights slice
		flight := JFlight{
			Code:   flightData.Code,
			From:   flightData.From,
			Time:   scheduledArrival,
			Status: status,
		}
		flights = append(flights, flight)
	}

	return flights, nil
}

// Function to print flight data
func PrintFlights(Flights []JFlight) {
	fmt.Println("Time From Code Status")
	for _, flight := range Flights {
		flightStatus := flight.GetStatus()
		formattedTime := flight.Time.Format("15:04")
		fmt.Printf("%s %s %s, %s\n", formattedTime, flight.From, flight.Code, flightStatus)
	}
}

func (f JFlight) GetStatus() string {
	if f.Status.Cancelled {
		return "Cancelled"
	} else if f.Status.Arrived.IsZero() {
		return f.Status.Expected.Format("15:04")
	} else {
		return f.Status.Arrived.Format("15:04")
	}

}

// val := reflect.ValueOf(j)

// 	for i := 0; i < val.NumField(); i++ {

//         fieldValue := val.Field(i)

// 		fieldType := fieldValue.Type()

// 		if fieldType.Kind() == reflect.String {
// 			time, error := time.Parse("15:04", j)

// }
// 	}
