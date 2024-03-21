package flights

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"reflect"
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
	Status JStatus   `json:"status"`
}

func (j JStatus) JStatusConvert() FStatus {
	val := reflect.ValueOf(j)

	for i := 0; i < val.NumField(); i++ {
        
        fieldValue := val.Field(i)

		fieldType := fieldValue.Type()

		if fieldType.Kind() == reflect.String {
			time, error := time.Parse("15:04", j)

}
	}

func (j JStatus) JStatusConvert() FStatus {
	val := reflect.ValueOf(j)
	fStatus := FStatus{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := field.Type()

		if fieldType.Kind() == reflect.String {
			timeStr := field.String()
			parsedTime, err := time.Parse("15:04", timeStr)
			if err != nil {
				fmt.Printf("Error parsing time string: %v\n", err)
				continue
			}

			switch i {
			case 0:
				fStatus.Arrived = parsedTime
			case 1:
				fStatus.Expected = parsedTime
			}
		}
	}

	fStatus.Cancelled = j.Cancelled
	return fStatus
}


func ReadJSONFile(filename string) ([]JFlight, error) {
	var data struct {
		Flights []JFlight `json:"flights"`
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

	// adds the json data to the private data struct to access the flight key for the nested flight objects
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Decode Json error: ", err)
		return nil, err
	}
	fmt.Println("Data:", data.Flights)
	return data.Flights, nil
}

// Function to print flight data
func PrintFlights(Flights []JFlight) {
	fmt.Println("Time From Code Status")
	for _, flight := range Flights {
		flightStatus := flight.GetStatus()
		fmt.Println("Flight status print", flightStatus)
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
