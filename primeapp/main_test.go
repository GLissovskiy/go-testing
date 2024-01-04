package main

import "testing"

// func Test_isPrime(t *testing.T) {
// 	result, msg := isPrime(0)
// 	if result {
// 		t.Errorf("with %d as test parameter, got true, but expected false", 0)
// 	}
// 	if msg != "0 is not prime, by definitions" {
// 		t.Error("wrong message returned:", msg)
// 	}

// 	result, msg = isPrime(7)
// 	if !result {
// 		t.Errorf("with %d as test parameter, got false, but expected true", 7)
// 	}
// 	if msg != "7 is a prime number!" {
// 		t.Error("wrong message returned:", msg)
// 	}
// }

func Test_isPrime(t *testing.T) {

	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
		{"zero", 0, false, "0 is not prime, by definitions"},
		{"one", 1, false, "1 is not prime, by definitions"},
		{"negative", -14, false, "negative numbers are not prime, by definitions."},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}

}
