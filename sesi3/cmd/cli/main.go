package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

type Student struct {
	ID      int
	Name    string
	Address string
	Job     string
	Reason  string
}

type StudentList []Student

var studentList = StudentList{
	Student{
		ID:      1,
		Name:    "Joshua",
		Address: "Jakarta",
		Job:     "Software Engineer",
		Reason:  "I want to learn Go",
	},
	Student{
		ID:      2,
		Name:    "Abdurrahman",
		Address: "Bandung",
		Job:     "Data Scientist",
		Reason:  "I want to learn Go",
	},
	Student{
		ID:      3,
		Name:    "Indra",
		Address: "Bogor",
		Job:     "DevOps Engineer",
		Reason:  "I want to learn Go",
	},
	{
		ID:      4,
		Name:    "Pramudya",
		Address: "Bekasi",
		Job:     "Backend Engineer",
		Reason:  "I want to learn Go",
	},
}

func linearSearch(studentList StudentList, studentNumber int) (Student, bool) {
	for _, student := range studentList {
		if studentNumber == student.ID {
			return student, true
		}
	}
	return Student{}, false
}

func binarySearch(studentList StudentList, studentNumber int) (Student, bool) {
	resultIndex, ok := slices.BinarySearchFunc(studentList, Student{studentNumber, "", "", "", ""}, func(student, student2 Student) int {
		return cmp.Compare(student.ID, student2.ID)
	})

	if !ok {
		log.Println("Student not found")
		return Student{}, false
	}

	result := studentList[resultIndex]

	return result, true
}

func main() {

	args := os.Args

	if len(args) < 2 {
		log.Println("Please provide a student number")
		return
	}

	fmt.Printf("%#v\n", args)

	studentNumber, err := strconv.Atoi(args[1])

	if err != nil {
		log.Println("Please provide a valid number", err)
		return
	}

	timeNow := time.Now()

	student, ok := linearSearch(studentList, studentNumber)

	if !ok {
		log.Println("Student not found")
		return
	}

	fmt.Println(student)

	fmt.Println("Linear search took", time.Since(timeNow))

	timeNow = time.Now()

	student, ok = binarySearch(studentList, studentNumber)

	if !ok {
		log.Println("Student not found")
		return
	}

	fmt.Println(student)

	fmt.Println("Binary search took", time.Since(timeNow))

}
