package main

import "testing"

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
