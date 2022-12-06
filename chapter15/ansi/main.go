package main

import (
	"fmt"
)

var fgColors = []string{
	"", "30", "31", "32", "33", "34", "35", "36", "37",
	"90", "91", "92", "93", "94", "95", "96", "97",
}

var bgColors = []string{
	"", "40", "41", "42", "43", "44", "45", "46", "47",
	"100", "101", "102", "103", "104", "105", "106", "107",
}

func main() {
	fmt.Printf("   ")
	for _, bg := range bgColors {
		fmt.Printf("%4s", bg)
	}
	fmt.Println()
	for _, fg := range fgColors {
		fmt.Printf("%2s ", fg)
		for _, bg := range bgColors {
			if len(fg) > 0 {
				if len(bg) > 0 {
					v := fmt.Sprintf("\x1b[%s;%sm Aa \x1b[0m", fg, bg)
					fmt.Print(v)
				} else {
					fmt.Printf("\x1b[%sm Aa \x1b[0m", fg)
				}
			} else if len(bg) > 0 {
				fmt.Printf("\x1b[%sm Aa \x1b[0m", bg)
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Println()
	}
}
