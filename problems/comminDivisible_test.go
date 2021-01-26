package main

import "testing"

/**
* Sample code for the test
 */
func TestGcdLenghth(t *testing.T) {
	// To test fail case
	testValue := gcd(5, 10)

	// To test pass case
	// testValue := gcd(0, 0)

	if testValue == 0 {
		t.Logf("The value is 0")
	} else {
		t.Errorf("The value is %v", testValue)
	}

}
