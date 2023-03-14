package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime"},
		{"not prime", 6, false, "6 is not a prime, it is divisible by 2"},
		{"negative", -1, false, "Negative numbers are not prime, by definition"},
		{"zero", 0, false, "0 is not a prime, by definition"},
		{"one", 1, false, "1 is not a prime, by definition"},
		{"negative par", -2, false, "Negative numbers are not prime, by definition"},
	}
	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected != result {
			t.Errorf("%s: expected %t, got %t", e.name, e.expected, result)
		}
		if e.msg != msg {
			t.Errorf("%s: expected %s, got %s", e.name, e.msg, msg)
		}

	}

}

func Test_prompt(t *testing.T) {
	//save the original stdout
	origStdout := os.Stdout

	//create a pipe to read from
	r, w, _ := os.Pipe()

	//set stdout to the pipe
	os.Stdout = w

	//call the function to test
	prompt()

	//close the pipe
	w.Close()

	//reset stdout
	os.Stdout = origStdout

	//read the output
	out, _ := io.ReadAll(r)

	//check the output
	if string(out) != "Enter a number: " {
		t.Errorf("expected 'Enter a number: ', got %s", string(out))
	}

}

func Test_intro(t *testing.T) {
	//save the original stdout
	origStdout := os.Stdout

	//create a pipe to read from
	r, w, _ := os.Pipe()

	//set stdout to the pipe
	os.Stdout = w

	//call the function to test
	intro()

	//close the pipe
	w.Close()

	//reset stdout
	os.Stdout = origStdout

	//read the output
	out, _ := io.ReadAll(r)

	//check the output
	if !strings.Contains(string(out), "Is it Prime?") {
		t.Errorf("expected 'Is it Prime?', got %s", string(out))
	}

}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"quit", "q", ""},
		{"not a number", "a", "Please enter a whole number!"},
		{"prime", "7", "7 is a prime"},
	}
	for _, e := range tests {
		//read a input from a string
		input := strings.NewReader(e.input)
		//create a scanner
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(e.expected, res) {
			t.Errorf("%s: expected '%s', got '%s'", e.name, e.expected, res)
		}

	}
}

func Test_readUserInput(t *testing.T) {
	//to test this function we need a channel, and an instance of an io.Reader
	doneChan := make(chan bool)

	// create a reference to  bytes.Buffer
	var stdin bytes.Buffer

	//set the input
	stdin.Write([]byte("1\nq\n"))

	//call the function
	go readUserInput(&stdin, doneChan)

	<-doneChan

	close(doneChan)

}
