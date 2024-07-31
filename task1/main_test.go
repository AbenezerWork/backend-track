package main

import (
	"testing"
)

func TestGetAverage(T *testing.T) {
	s := Student{name: "test1", subjects: []Subject{{name: "math", grade: 10}, {name: "english", grade: 20}, {name: "SAT", grade: 30}}}
	if s.GetAverage() != 20 {
		T.Error("s.GetAverage = 20")
	}
}
