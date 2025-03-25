package main

import "fmt"

func grade_calculator(grades map[string]int) int {
	total := 0
	if len(grades) == 0 {
		return 0
	}
	for _, grade := range grades {
		total += grade
	}
	return total / len(grades)
}

func accept_info() (string, map[string]int) {
	grades := make(map[string]int)
	var name string
	var num_subjects int

	fmt.Println("Name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter the number of subjects: ")
	if num_subjects < 0 {
		fmt.Println("Invalid number of subjects. Please enter a positive integer.")
		return "", nil
	}
	fmt.Scanln(&num_subjects)

	for i := 0; i < num_subjects; i++ {
		var grade int
		var subject string
		fmt.Println("Enter the subject for subject", i+1, ": ")
		fmt.Scanln(&subject)
		fmt.Println("Enter the grade for subject", i+1, ": ")
		fmt.Scanln(&grade)
		if grade < 0 || grade > 100 {
			fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
			i--
			continue
		}
		grades[subject] = grade
	}
	return name, grades
}

func main() {
	name, grades := accept_info()

	average := grade_calculator(grades)
	fmt.Printf("Student: %s\n", name)
	for subject, grade := range grades {
		fmt.Printf("Subject %s: %d\n", subject, grade)
	}
	fmt.Printf("Average grade: %d\n", average)
}
