package main

import (
  "fmt"
  "os"
  "regexp"
  "strings"
  "strconv"
)

func main() {
	file, err := os.ReadFile("input.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }
    input := string(file) // convert content to a 'string'

	r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	instructions := r.FindAllString(input, -1)

	var result float64
	for i := 0; i < len(instructions); i++ {
		values := strings.Split(strings.Replace(strings.Replace(instructions[i], "mul(", "", -1), ")", "", -1), ",")
		var left float64
		var right float64
		if s, err := strconv.ParseFloat(values[0], 64); err == nil {
			left = s
		  }
		  if s, err := strconv.ParseFloat(values[1], 64); err == nil {
			right = s
		  }
		result += left*right
	}
	fmt.Printf("Sum: %f\n", result)
  }