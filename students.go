package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func help() {
	fmt.Println("Grading system")
	fmt.Println("add-student [name] - Add student")
	fmt.Println("select-student [name] - Select a student")
	fmt.Println("add-grade [course] [grade] - Add a grade for a course for the current selected student")
	fmt.Println("end - Stop the program")
}

type Course struct {
	name   string
	grades *list.List
}

type Student struct {
	name    string
	courses *list.List
}

func (student Student) Totals() {
	fmt.Println("Grades for " + student.name)

	for c := student.courses.Front(); c != nil; c = c.Next() {
		var course = c.Value.(Course)
		course.Totals()
	}
}

func (course Course) Totals() {
	var total float32 = 0
	for n := course.grades.Front(); n != nil; n = n.Next() {
		total += n.Value.(float32)
	}
	var avg float32 = total / float32(course.grades.Len())

	fmt.Printf("%s - %d exams - %3.2f avg", course.name, course.grades.Len(), avg)
	fmt.Println("")
}

func ToFloat(str string) float32 {
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0.0
	}
	return float32(value)
}

func main() {
	var students = list.New()
	var student Student

	for {
		fmt.Print("Input command:")
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Please enter a valid command", err)
			help()
			continue
		}
		// remove the delimeter from the string
		command = strings.TrimSuffix(command, "\n")

		if command == "help" {
			help()
			continue
		}
		if command == "end" || command == "quit" {
			break
		}
		if strings.HasPrefix(command, "add-student") {
			var name = strings.Split(command, " ")[1]
			student = Student{name, list.New()}
			students.PushBack(student)
		}
		if strings.HasPrefix(command, "select-student") {
			var name = strings.Split(command, " ")[1]
			var found = false
			for s := students.Front(); s != nil; s = s.Next() {
				if s.Value.(Student).name == name {
					student = s.Value.(Student)
					found = true
				}
			}
			if found {
				fmt.Printf("Student %s is selected", student.name)
				fmt.Println("")
			} else {
				fmt.Printf("Student %s not found", name)
				fmt.Println("")
			}
		}
		if strings.HasPrefix(command, "add-grade") {
			var parts = strings.Split(command, " ")
			var courseName = parts[1]
			var grade = parts[2]

			var found = false
			for s := student.courses.Front(); s != nil; s = s.Next() {
				if s.Value.(Course).name == courseName {
					var course = s.Value.(Course)
					course.grades.PushBack(ToFloat(grade))
					found = true
				}
			}
			if !found {
				var course = Course{courseName, list.New()}
				course.grades.PushBack(ToFloat(grade))
				student.courses.PushBack(course)
			}
		}
	}

	for s := students.Front(); s != nil; s = s.Next() {
		var student = s.Value.(Student)
		student.Totals()
	}
}
