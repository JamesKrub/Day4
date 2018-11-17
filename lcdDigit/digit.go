package lcdDigit

import (
	"fmt"
	"math"
)

// type digits struct {
// 	Num string
// }

type digits struct {
	Top string
	Mid string
	Bot string
}

// MakeDigit is creating a digit picture
func MakeDigit(n int) {
	// 	draws := []digits{
	// 		{` _
	// | |
	// |_|
	// `},
	// 		{`
	// |
	// |
	// `},
	// 		{` _
	//  _|
	// |_
	// `},
	// 		{` _
	// _|
	// _|
	// `},
	// 		{`
	// |_|
	//   |
	// `},
	// 		{` _
	// |_
	//  _|
	// `},
	// 		{` _
	// |_
	// |_|
	// `},
	// 		{` _
	// |
	// |
	// `},
	// 		{` _
	// |_|
	// |_|
	// `},
	// 		{` _
	// |_|
	//  _|
	// `},
	// 	}

	draws := []digits{
		{" _ ", "| |", "|_|"},
		{"   ", "  |", "  |"},
		{" _ ", " _|", "|_ "},
		{" _ ", " _|", " _|"},
		{"   ", "|_|", "  |"},
		{" _ ", "|_ ", " _|"},
		{" _ ", "|_ ", "|_|"},
		{" _ ", "  |", "  |"},
		{" _ ", "|_|", "|_|"},
		{" _ ", "|_|", "  |"},
	}

	var count, i int = 0, n
	digit := []int{}
	for i != 0 {
		i /= 10
		count = count + 1
	}
	for i := 1; i <= count; i++ {

		if i == 1 {
			store := n % 10
			digit = append(digit, store)

		} else {
			store := (n / (int(math.Pow(10, (float64(i) - 1))))) % 10
			digit = append(digit, store)
		}
	}

	for dis := count; dis > 0; dis-- {
		// fmt.Println(digit[count-1])
		fmt.Printf(draws[digit[dis-1]].Top)
	}
	fmt.Printf("\n")

	for dis := count; dis > 0; dis-- {
		fmt.Printf(draws[digit[dis-1]].Mid)
	}
	fmt.Printf("\n")

	for dis := count; dis > 0; dis-- {
		fmt.Printf(draws[digit[dis-1]].Bot)
	}
	fmt.Printf("\n\n")

}
