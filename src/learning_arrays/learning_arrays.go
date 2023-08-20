package main

import "fmt"

func main() {
	names := make([]string, 4)
	names[0] = "Keith"
	names[1] = "Kevin"
	names[2] = "Kayla"
	names[3] = "Kyle"

	fmt.Println(names)
	fmt.Println("names[2] is nil", names[2] == "")
}
