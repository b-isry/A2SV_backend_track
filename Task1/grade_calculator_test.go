package main

import (
	"testing"
)

func TestGradeCalculator(t *testing.T) {
	grades := map[string]int{
		"Math": 90,
		"English": 85,
		"Science": 95,
	}
	expected_average := 90
	average := grade_calculator(grades)
	if average != expected_average {
		t.Errorf("Expected average grade to be %d, but got %d", expected_average, average)
	}
	
	
}
func TestGradeCalculatorEmpty(t *testing.T) {
	grades := map[string]int{
	}
	expected_average := 0
	average := grade_calculator(grades)
	if average != expected_average {
		t.Errorf("Expected average grade to be %d, but got %d", expected_average, average)
	}
}

