package main

import "fmt"

func main() {
	birthdays := map[string]string{
		"Keith": "02/06/1990",
		"Kevin": "01/01/1958",
		"Kayla": "06/24/1986",
	}
	fmt.Println(birthdays)

	ages := map[string]int{}
	ages["Keith"] = 29
	ages["Kevin"] = 71
	ages["Kayla"] = 43
	ages["Keith"] = 1

	fmt.Println(ages["Keith"])
}
