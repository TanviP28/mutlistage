package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func calculateAge(birthdate time.Time) int {
	currentTime := time.Now()
	years := currentTime.Year() - birthdate.Year()

	// Adjust age if the current date is before the birthday
	if currentTime.YearDay() < birthdate.YearDay() {
		years--
	}

	return years
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Remove leading/trailing whitespaces

	fmt.Print("Enter your birthdate (YYYY-MM-DD): ")
	birthdateInput, _ := reader.ReadString('\n')
	birthdateInput = strings.TrimSpace(birthdateInput) // Remove leading/trailing whitespaces

	birthdate, err := time.Parse("2006-01-02", birthdateInput)
	if err != nil {
		fmt.Println("Invalid birthdate. Please enter a valid date in the format YYYY-MM-DD.")
		return
	}

	age := calculateAge(birthdate)

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
