package flights

import (
	"fmt"
	"sort"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello from the flights package!")
}

// creates flights
type Flight struct {
	Code    string    // a string representing the flight code, eg. BA 134
	Origin  string    // a string representing the name of the airport where the flights came from
	DueTime time.Time // a time.Time representing the time that it is due.

}

// Creates a flight board and adds flights into a map
type Board struct {
	Name    string
	Flights map[string]Flight
}

// prints flights from flight struct
func ToString(f Flight) string {
	dueTimeString := f.DueTime.Format("15:04")
	return fmt.Sprintf("Flight %s from %s is expected at %s", f.Code, f.Origin, dueTimeString)

}

// Prints the display board struct
func (b Board) PrintBoard() {
	fmt.Println(b.Name, "\nTime From Code")
	for _, Flight := range b.Flights {
		dueTimeString := Flight.DueTime.Format("15:04")
		fmt.Printf("%s %s %s\n", dueTimeString, Flight.Origin, Flight.Code)
	}
	fmt.Println()
}

// sorts the flights into time order and prints them formatted
func (b Board) Display() {

	var flights []Flight
	for _, flight := range b.Flights {
		flights = append(flights, flight)
	}

	sort.Slice(flights, func(i, j int) bool {
		return flights[i].DueTime.Before(flights[j].DueTime)
	})

	fmt.Println("Time From Code")
	for _, flight := range flights {
		dueTimeString := flight.DueTime.Format("15:04")
		fmt.Printf("%s %s %s\n", dueTimeString, flight.Origin, flight.Code)
	}
}
