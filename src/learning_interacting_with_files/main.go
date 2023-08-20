package main

import (
	"bufio"
	"fmt"
	"motd/message"
	"os"
	"strings"
	//"io/ioutil"
)

func main() {
	f, err := os.OpenFile("C:\\Code_Workspace\\Golang\\src\\learning_interacting_with_files\\text.txt", os.O_WRONLY, 0644)
	
	// Will show an error if the file doesn't exist
	if err != nil {
		fmt.Println("Error: Unable to open /etc/motd")
		os.Exit(1)
	}

	// Will close the file after everything is done running
	defer.f.Close()
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	message := message.Greeting(name, phrase)
	_, err = f.Write([]byte(message))
	
	//err := ioutil.WriteFile("C:\\Code_Workspace\\Golang\\src\\learning_interacting_with_files\\text.txt", []byte(message), 0644)
		
	if err != nil {
		fmt.Println("Error: Filed to write to /etc/motd")
		os.Exit(1)
	}
}
