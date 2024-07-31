package main

import (
	"testing"
)

func TestFreqCount(T *testing.T) {

	returned := FreqCount("hello, world. Hello. World!!")
	expected := map[string]int{"hello": 2, "world": 2}
	isequal := true
	for key := range expected {
		if returned[key] != expected[key] {
			isequal = false
			break
		}
	}
	if !isequal {
		T.Error("returned is not equal to expected! returned: ", returned, " expected: ", expected)
	}

}

func TestPalindromeCheck(T *testing.T) {
	if false == (PalindromeCheck("race car") && PalindromeCheck("Race Car") && PalindromeCheck("Race, Car")) {
		T.Error("one of the following is not correct \"race car\"  \"Race Car\" \"Race, Car\"")
	}
}
