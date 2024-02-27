package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name  string
	Age   int8
	Notes []string
}

func add(students []Student) {
	var name string
	var age int8
	var str_notes string
	fmt.Println("Enter student name: ")
	fmt.Scan(&name)
	fmt.Println("Enter student age:")
	fmt.Scan(&age)
	fmt.Println("Enter student notes separated by comma:")
	fmt.Scan(&str_notes)
	notes := strings.Split(str_notes, ",")
	students = append(students, Student{name, age, notes})
}

func print(students []Student) {
	fmt.Println(students)
}

func main() {
	var students []Student
	var command string
	for {
		fmt.Println("Type \"add\" for adding a student, \"print\" for printing info about all students or \"end\" to stop the program")
		// fmt.Println(command)
		fmt.Scan(&command)
		if command == "add" {
			add(students)
			continue
		}
		if command == "print" {
			print(students)
			continue
		}
		if command == "end" {
			break
		}
	}
	fmt.Println("Program ended")
}
