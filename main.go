package main

import "fmt"

func gradingStudents(grades []int) []int {
	result := make([]int, len(grades))

	for i, grade := range grades {
		if grade < 38 {
			result[i] = grade
		} else {
			nextMultiple := (grade + 4) / 5 * 5
			if nextMultiple-grade < 3 {
				result[i] = nextMultiple
			} else {
				result[i] = grade
			}
		}
	}

	return result
}

func main() {
	grades := []int{73, 67, 38, 33}
	fmt.Println(gradingStudents(grades))
}
