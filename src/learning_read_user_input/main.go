package main

import (
	"bufio"
	"fmt"
	"motd/message"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	message := message.Greeting(name, phrase)
	err := ioutil.WriteFile("/etc/motd", []byte(message), 0644)	

	if err != nil {
		fmt.Println("Unable to write /etc/motd")
		os.Exit(1)
	}
}
