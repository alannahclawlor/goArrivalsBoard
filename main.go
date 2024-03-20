package main

import (
	"fmt"
	"sort"

	"golang.makers.tech/go_arrivals_project/flights"
)

func main() {
	fmt.Println("Welcome to the Arrivals Board!")

	// creates the arrivals board and adds the json fights to the Jflight struct
	ArrivalsBoard, err := flights.ReadJSONFile("flights/flightData.json")
	if err != nil {
		println(err.Error())
	}

	// sorts the flights by time order (not sure if this should be in the json.go code)
	sort.Slice(ArrivalsBoard, func(i, j int) bool {
		return ArrivalsBoard[i].Time.Before(ArrivalsBoard[j].Time)
	})

	// returns the arrivals board
	flights.PrintFlights(ArrivalsBoard)

}

// PREVIOUS CODE

// fmt.Println("Welcome to the Arrivals Board!")
// flights.HelloWorld() // Using the public function
// Flight1 := flights.Flight{
// 	Code:    "BA 341",
// 	Origin:  "London",
// 	DueTime: time.Date(2024, time.March, 30, 18, 0, 0, 0, time.Local),
// }
// Flight2 := flights.Flight{
// 	Code:    "LH 712",
// 	Origin:  "Frankfurt",
// 	DueTime: time.Date(2024, time.March, 30, 15, 30, 0, 0, time.Local),
// }
// Flight3 := flights.Flight{
// 	Code:    "AF 123",
// 	Origin:  "Paris",
// 	DueTime: time.Date(2024, time.March, 30, 13, 45, 0, 0, time.Local),
// }

// // fmt.Println(flights.ToString(Flight1))
// // fmt.Println(flights.ToString(Flight2))
// // fmt.Println(flights.ToString(Flight3))

// board := flights.Board{
// 	Name: "Inbound Flights",
// 	Flights: map[string]flights.Flight{
// 		"flight1": Flight1,
// 		"flight2": Flight2,
// 		"flight3": Flight3,
// 	},
// }

// // board.PrintBoard()
// board.Display()
