# Grading Students

This Go program implements a function `gradingStudents` that takes a slice of student grades and returns a new slice with the grades rounded up to the nearest multiple of 5, if the grade is at least 38 and the rounding does not result in a grade greater than 100.

## Function Documentation

```go
func gradingStudents(grades []int) []int
```

The `gradingStudents` function takes a slice of integer grades as input and returns a new slice of integer grades.

### Parameters

- `grades []int`: A slice of integer grades, where each grade is a value between 0 and 100 (inclusive).

### Return Value

- `[]int`: A new slice of integer grades, where the grades have been rounded up to the nearest multiple of 5 if the grade is at least 38 and the rounding does not result in a grade greater than 100.

### Algorithm

1. Create a new slice `result` with the same length as the `grades` slice.
2. Iterate through each grade in the `grades` slice.
3. If the grade is less than 38, assign the original grade to the corresponding index in the `result` slice.
4. If the grade is 38 or greater, calculate the next multiple of 5 that is greater than or equal to the grade.
5. If the difference between the next multiple and the original grade is less than 3, assign the next multiple to the corresponding index in the `result` slice.
6. If the difference between the next multiple and the original grade is 3 or greater, assign the original grade to the corresponding index in the `result` slice.
7. Return the `result` slice.

## Example Usage

```go
func main() {
    grades := []int{73, 67, 38, 33}
    fmt.Println(gradingStudents(grades))
}

```

```
[75 67 40 33]

```

In this example, the grade 73 is rounded up to 75, the grade 67 is not rounded, the grade 38 is rounded up to 40, and the grade 33 is not rounded.
