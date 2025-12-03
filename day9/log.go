package day9

import (
	"log"
	"os"
)

func Logs() {

	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal("cannot open:", err)
	}
	log.SetOutput(file)
	log.Println("This is a log message")
}
