package main

import (
	"Project1/wordfreq"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please input any length of text:")
	in := bufio.NewReader(os.Stdin)
	userInput, _ := in.ReadString('\n')
	jsonOutput := wordfreq.WordCountService(userInput)
	fmt.Println(string(jsonOutput))
	//
}
