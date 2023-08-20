package main

import (
	"fmt"
	"motd/message"
)

func main() {
	message := message.Greeting("Keith", "Hello")
	fmt.Println(message)
}
