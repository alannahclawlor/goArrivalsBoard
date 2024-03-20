package flights

import (
	"testing"
	"time"

	"fmt"
	"strings"

	"golang.makers.tech/go_arrivals_project/test_utils"
)

// tests for the individual flights
func TestFlightOne(t *testing.T) {
	Flight1 := Flight{
		Code:    "BA 341",
		Origin:  "London",
		DueTime: time.Date(2024, time.March, 30, 14, 0, 0, 0, time.Local),
	}

	result, expected := ToString(Flight1), "Flight BA 341 from London is expected at 14:00"
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestFlightTwo(t *testing.T) {
	Flight2 := Flight{
		Code:    "LH 712",
		Origin:  "Frankfurt",
		DueTime: time.Date(2024, time.March, 30, 15, 30, 0, 0, time.Local),
	}

	result, expected := ToString(Flight2), "Flight LH 712 from Frankfurt is expected at 15:30"
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestFlightThree(t *testing.T) {
	Flight3 := Flight{
		Code:    "AF 123",
		Origin:  "Paris",
		DueTime: time.Date(2024, time.March, 30, 13, 45, 0, 0, time.Local),
	}
	result, expected := ToString(Flight3), "Flight AF 123 from Paris is expected at 13:45"
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

// test for the displayboard struct
func TestDisplayBoard(t *testing.T) {
	Flight1 := Flight{
		Code:    "BA 341",
		Origin:  "London",
		DueTime: time.Date(2024, time.March, 30, 14, 0, 0, 0, time.Local),
	}
	Flight2 := Flight{
		Code:    "LH 712",
		Origin:  "Frankfurt",
		DueTime: time.Date(2024, time.March, 30, 15, 30, 0, 0, time.Local),
	}
	Flight3 := Flight{
		Code:    "AF 123",
		Origin:  "Paris",
		DueTime: time.Date(2024, time.March, 30, 13, 45, 0, 0, time.Local),
	}

	board := Board{
		Name: "Inbound Flights",
		Flights: map[string]Flight{
			"flight1": Flight1,
			"flight2": Flight2,
			"flight3": Flight3,
		},
	}

	rec := test_utils.StartRecording()

	// Call the Display method
	board.Display()

	result := test_utils.EndRecording(rec)

	expected := `Time From Code
13:45 Paris AF 123
14:00 London BA 341
15:30 Frankfurt LH 712`

	// Compare result with expected
	if result != expected {
		t.Errorf("Result is\n%s\nbut expected\n%s\n", result, expected)
	}
}

func TestReadJSONFile(t *testing.T) {
	// Read flight data from JSON file
	flightData, err := ReadJSONFile("flightData.json")
	if err != nil {
		t.Fatalf("Error reading JSON file: %v", err)
	}

	// Format flight data into string
	var builder strings.Builder
	builder.WriteString("Time From Code\n")
	for _, flight := range flightData {
		formattedTime := flight.Time.Format("15:04")
		builder.WriteString(fmt.Sprintf("%s %s %s\n", formattedTime, flight.From, flight.Code))
	}
	expected := builder.String()

	rec := test_utils.StartRecording()

	// Call the Display method
	PrintFlights(flightData)

	result := test_utils.EndRecording(rec)

	// trim whitespace so the tests pass!
	result = strings.TrimSpace(result)
	expected = strings.TrimSpace(expected)

	// Compare result with expected
	if result != expected {
		t.Errorf("Result is\n%s\nbut expected\n%s\n", result, expected)
	}
}
