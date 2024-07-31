package main

import (
	"fmt"
)

type Subject struct {
	name  string
	grade float32
}

type Student struct {
	name     string
	subjects []Subject
}

func (s Student) GetAverage() float32 {
	var tot float32
	cnt := float32(len(s.subjects))
	for _, subject := range s.subjects {
		tot += subject.grade
	}

	return tot / cnt
}

func getString() string {
	var ret string
	fmt.Scan(&ret)
	for len(ret) <= 2 || len(ret) > 20 {
		fmt.Println("You must enter a name of length between 3 and 20!!")
		fmt.Println("Try again:")
		fmt.Scan(&ret)
	}
	return ret
}

func getInt() int {
	var n int
	fmt.Scan(&n)
	for n <= 1 || n > 20 {
		fmt.Println("You must enter a number between 1 and 20!!")
		fmt.Println("Try again:")
		fmt.Scan(&n)
	}
	return n

}
func getFloat() float32 {
	var ret float32
	fmt.Scan(&ret)
	if ret < 0 || ret > 100 {
		fmt.Println("Please enter a valid grade!!")
	}
	return ret
}

func main() {
	fmt.Println("Input your name:")
	name := getString()
	student := Student{name: name}

	fmt.Println("Hi ", name, " how many subjects do you have?")
	n := getInt()

	fmt.Println("For every subject input the name and grade respectively.")
	subjects := make([]Subject, n)
	i := 0

	for i < n {
		fmt.Printf("Subject %v: ", i)
		subjects[i].name = getString()
		fmt.Scan(&subjects[i].grade)
		i++
		fmt.Println(max(67, 21))
	}
	student.subjects = subjects
	fmt.Printf("Your average grade is %v. \n", student.GetAverage())
}
