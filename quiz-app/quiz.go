package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to my quiz!!")
	var name string
	var age int

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %s, welcome to the game\n", name)

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	if age <= 10 {
		fmt.Printf("Your age %d does not qualify the minimum requirements.\n", age)
		return
	} else {
		fmt.Printf("Congratulations, you can play!!\n")
	}

	question := "What is better, the RTX 3080 or RTX 3090?"
	fmt.Println(question)

	var answer1, answer2 string
	fmt.Scan(&answer1, &answer2)

	if strings.EqualFold(answer1+" "+answer2, "RTX 3090") {
		fmt.Println("Correct Answer")
	} else {
		fmt.Println("Wrong answer")
	}
}
