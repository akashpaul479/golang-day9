package day9

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type students3 struct {
	Name  string
	Age   int
	Grade string
}

var student []students3
var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		os.Exit(1)
	}
	logger = log.New(logFile, "logfile:", log.Ldate|log.Ltime|log.Lshortfile)

}
func saveStudents(filename string) error {
	data, err := json.MarshalIndent(student, "", " ")
	if err != nil {
		logger.Println("Error: Failed to marshal json", err)
		return err
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		logger.Printf("Error failed to saved file %s: %v", filename, err)
		return err
	}
	logger.Printf("success: saved %d students to be %s", len(student), filename)
	return nil

}

func Loadstudents(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		logger.Println("Error: failed to load file:", filename, err)
		return err
	}
	if err := json.Unmarshal(data, &student); err != nil {
		logger.Printf("Error: failed to unmarshal json %v", err)
		return err
	}
	logger.Printf("Success: loaded %d from students %s", len(student), filename)
	return nil
}
func Addstudent(name string, age int, grade string) {
	st := students3{Name: name, Age: age, Grade: grade}
	student = append(student, st)

	logger.Printf("Info: student added ->Name:%s Age:%d Grade:%s", name, age, grade)
}
func Liststudents() {
	logger.Printf("Info: listing all %d students", len(student))

	for _, s := range student {
		fmt.Printf("Name:%s , Age:%d , Grade:%s\n", s.Name, s.Age, s.Grade)
	}
}
func Projects3() {
	Loadstudents("students.json")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("__students database")
		fmt.Println("1.Add students")
		fmt.Println("2.List students")
		fmt.Println("3.save and exit")

		fmt.Println("Enter choice:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Printf("Enter name:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Printf("Enter age:")
			ages, _ := reader.ReadString('\n')
			ages = strings.TrimSpace(ages)
			age, _ := strconv.Atoi(ages)

			fmt.Printf("Enter grade:")
			grade, _ := reader.ReadString('\n')
			grade = strings.TrimSpace(grade)

			Addstudent(name, age, grade)
			fmt.Println("Students added succesfully.")
			logger.Printf("Succes students %s Added:", name)

		case "2":
			fmt.Println("All students")
			Liststudents()

		case "4":
			saveStudents("students.json")
			fmt.Println("saved succesfully")
			logger.Println("program excited by user")
			return
		default:
			fmt.Println("Invalid choice please enter again.")
			logger.Println("Warning:user entered invalid choice :", input)

		}
	}

}
