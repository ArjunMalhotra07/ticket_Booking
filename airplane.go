package main

import (
	"fmt"
)

func main() {
	f := fmt.Println
	p := fmt.Print
	s := fmt.Scan

	places := map[string]int{
		"Atlanta":      0,
		"Boston":       1,
		"Chicago":      2,
		"Miami":        3,
		"Philadelphia": 4,
	}
	var destination [5][5]int
	destination[0][0] = 0
	destination[0][1] = 0
	destination[0][2] = 0
	destination[0][3] = 715
	destination[0][4] = 0

	destination[1][0] = 0
	destination[1][1] = 0
	destination[1][2] = 701
	destination[1][3] = 0
	destination[1][4] = 711

	destination[2][0] = 0
	destination[2][1] = 702
	destination[2][2] = 0
	destination[2][3] = 708
	destination[2][4] = 0

	destination[3][0] = 0
	destination[3][1] = 0
	destination[3][2] = 0
	destination[3][3] = 0
	destination[3][4] = 718

	destination[4][0] = 713
	destination[4][1] = 712
	destination[4][2] = 705
	destination[4][3] = 717
	destination[4][4] = 0
	for {

		f()
		f("************************************************************")
		f("* 	Places: Atlanta, Boston, Chicago, Miami, Philadelphia.")
		p("* 	Enter Origin --- ")
		var origin string
		s(&origin)
		p("* 	Enter Destination --- ")
		var dest string
		s(&dest)
		origin_Code := places[origin]
		destination_Code := places[dest]
		flight_Number := destination[origin_Code][destination_Code]
		if flight_Number == 0 {
			f("*	No flights Available Currently. Try Again Later")
		} else {
			f("*	Flight Number ", flight_Number)
		}
	}

}
