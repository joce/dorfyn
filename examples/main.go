package main

import (
	"fmt"

	"github.com/joce/dorfyn"
)

// This file lists several usage examples of this library and can be used to verify behavior.
func main() {

	// Basic quote example.
	// --------------------
	{
		q, err := dorfyn.GetQuotes([]string{"^DJI", "AAPL", "BTC-USD", "CADUSD=X", "CL=F", "FMAGX", "VT"})

		if err != nil {
			fmt.Println(err)
		} else {
			for _, v := range q {
				fmt.Printf("%s: %f\n", *v.Symbol, *v.RegularMarketPrice)
			}
		}
		fmt.Println()
	}
}
