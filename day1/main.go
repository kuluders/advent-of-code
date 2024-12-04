package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"slices"
	"math"
	"strconv"
)

func main() {
	// Step 1: Open the file
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Step 2: Create a reader
	r := bufio.NewReader(f)

	var left_list []string
	var right_list []string

	// Step 3: Read line by line
	for {
			line, err := r.ReadString('\n')
			if err != nil {
					break
			}
			pair := strings.Fields(line)
			left_list = append(left_list, pair[0])
			right_list = append(right_list, pair[1])
	}
	
	slices.Sort(left_list)
	slices.Sort(right_list)
	// fmt.Println(left_list, right_list)
	
	sum := 0.0
	extended_sum := 0.0
	type repeats struct {
		num float64
		count int
	}

	var occurrences []repeats
	var repeated []float64
	for i := 0; i < len(left_list); i++ {
		left := 0.0
		right := 0.0
		if s, err := strconv.ParseFloat(left_list[i], 64); err == nil {
			left = s
		}
		if s, err := strconv.ParseFloat(right_list[i], 64); err == nil {
			right = s
		}
		if slices.Contains(repeated, right) {
			for j := 0; j < len(occurrences); j++ {
				if occurrences[j].num == right {
					occurrences[j].count += 1
				}
			}
		} else {
			repeated = append(repeated, right)
			occurrences = append(occurrences, repeats{right, 1})
		}
		current_sum := math.Abs(left - right)
		fmt.Println(left, right, current_sum)
		sum += math.Abs(left-right)
	}

	fmt.Println(sum)

	for i := 0; i < len(left_list); i++ {
		left := 0.0
		if s, err := strconv.ParseFloat(left_list[i], 64); err == nil {
			left = s
		}
		if slices.Contains(repeated, left) {
			for j := 0; j < len(occurrences); j++ {
				if occurrences[j].num == left {
					extended_sum += occurrences[j].num * float64(occurrences[j].count)
				}
			}
		}
	}

	fmt.Printf("Sum: %f\n", extended_sum)
	fmt.Printf("Sum: %f\n", sum)
fmt.Println(sum)
	// Step 4: Close the file
	defer f.Close()
}
