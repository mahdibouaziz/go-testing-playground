package main

import (
	"bufio"
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
		{name: "prime", testNum: 7, expected: true, msg: "7 is a prime number!"},
		{name: "not prime", testNum: 8, expected: false, msg: "8 is not prime, it is divisible by 2!"},
		{name: "zero", testNum: 0, expected: false, msg: "0 is not prime, by definition!"},
		{name: "one", testNum: 1, expected: false, msg: "1 is not prime, by definition!"},
		{name: "negative", testNum: -7, expected: false, msg: "Negative numbers are not prime, by definition!"},
	}

	for _, primeTest := range primeTests {
		result, msg := isPrime(primeTest.testNum)
		if result != primeTest.expected {
			t.Errorf("%s: with %d as test param, got '%t', but expected '%t'", primeTest.name, primeTest.testNum, result, primeTest.expected)
		}
		if msg != primeTest.msg {
			t.Errorf("%s: wrong message returned, got '%s', but expected '%s'", primeTest.name, msg, primeTest.msg)
		}
	}

}

func Test_prompt(t *testing.T) {
	//  save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe, so when we call prompt,
	// the result will be on the reader of the pipe, in our case`r`
	os.Stdout = w
	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of prompt() from the read pipe
	out, _ := io.ReadAll(r)

	// perform our tests
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected '-> ' but got %s", string(out))
	}

}

func Test_intro(t *testing.T) {
	//  save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe, so when we call prompt,
	// the result will be on the reader of the pipe, in our case`r`
	os.Stdout = w
	intro()

	// close our writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of prompt() from the read pipe
	out, _ := io.ReadAll(r)

	// perform our tests
	expectedText := "Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit."
	if !strings.Contains(string(out), expectedText) {
		t.Errorf("incorrect prompt: expected '%s' but got %s", expectedText, string(out))
	}

}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "negative", input: "-7", expected: "Negative numbers are not prime, by definition!"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		reader := bufio.NewScanner(input)

		res, _ := checkNumbers(reader)
		if !strings.EqualFold(res, test.expected) {
			t.Errorf("%s: incorrect value returned, got '%s', expected '%s'", test.name, res, test.expected)
		}
	}

}
