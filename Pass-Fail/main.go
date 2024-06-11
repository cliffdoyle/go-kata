// Pass-Fail is a program that reads grade from the keyboard(standard input)
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	return number, nil
}

func main() {
	var count [3]int
	count[0]++
	count[2]++
	count[2]++

	count1:=fmt.Sprintf("This the array increment in action %#v",count)

	fmt.Println(count1)


	fmt.Print("Enter grade: ")
	grade, err := getFloat()//Call getFloat to get a grade
	if err != nil {
		log.Fatal(err)
	}
	status := ""

	if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}

	fmt.Println("A grade of", grade, "is", status) 

	
}
