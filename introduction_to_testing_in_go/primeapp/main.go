package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	//print a welcome message
	intro()

	//create a channel to indicate when the usar wants to quit
	doneChan := make(chan bool)

	//create a channel to receive the user input
	go readUserInput(os.Stdin, doneChan)

	//block until the user wants to quit
	<-doneChan

	// close the channel
	close(doneChan)

	//print a goodbye message
	fmt.Println("Goodbye.")

}

func readUserInput(in io.Reader, doneChan chan bool) {

	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	//read the user input
	scanner.Scan()

	//check if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?\n------------------\nEnter a number to see if it is a prime number.\nEnter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("Enter a number: ")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime, by definition", n)
	}
	//negative numbers
	if n < 0 {
		return false, fmt.Sprintf("Negative numbers are not prime, by definition")
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime, it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime", n)

}
