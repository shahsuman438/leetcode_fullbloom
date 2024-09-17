package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/shahsuman438/leetcode_fullbloom/pkg/logic"
)

func TestEarliestFullBloom1(t *testing.T) {
	plantTime := []int{1, 4, 3}
	growTime := []int{2, 3, 1}
	expected := 9
	start := time.Now()

	result := logic.EarliestFullBloom(plantTime, growTime)

	elapsed := time.Since(start)
	fmt.Printf("TestEarliestFullBloom1 Execution time: %s\n", elapsed)

	// Check the result
	if result != expected {
		t.Errorf("For plantTime %v and growTime %v, expected %d but got %d",
			plantTime, growTime, expected, result)
	}
}

func TestEarliestFullBloom2(t *testing.T) {
	plantTime := []int{1, 2, 3, 2}
	growTime := []int{2, 1, 2, 1}
	expected := 9
	start := time.Now()

	result := logic.EarliestFullBloom(plantTime, growTime)

	elapsed := time.Since(start)
	fmt.Printf("TestEarliestFullBloom2 Execution time: %s\n", elapsed)

	// Check the result
	if result != expected {
		t.Errorf("For plantTime %v and growTime %v, expected %d but got %d",
			plantTime, growTime, expected, result)
	}
}
func TestEarliestFullBloom3(t *testing.T) {
	plantTime := []int{1}
	growTime := []int{1}
	expected := 2
	start := time.Now()

	result := logic.EarliestFullBloom(plantTime, growTime)

	elapsed := time.Since(start)
	fmt.Printf("TestEarliestFullBloom3 Execution time: %s\n", elapsed)

	// Check the result
	if result != expected {
		t.Errorf("For plantTime %v and growTime %v, expected %d but got %d",
			plantTime, growTime, expected, result)
	}
}
func TestEarliestFullBloom4(t *testing.T) {
	plantTime := []int{3, 2, 4, 1}
	growTime := []int{5, 1, 6, 4}
	expected := 12
	start := time.Now()

	result := logic.EarliestFullBloom(plantTime, growTime)

	elapsed := time.Since(start)
	fmt.Printf("TestEarliestFullBloom4 Execution time: %s\n", elapsed)

	// Check the result
	if result != expected {
		t.Errorf("For plantTime %v and growTime %v, expected %d but got %d",
			plantTime, growTime, expected, result)
	}
}
func TestEarliestFullBloom5(t *testing.T) {
	plantTime := []int{2, 2, 2}
	growTime := []int{3, 3, 3}
	expected := 9
	start := time.Now()

	result := logic.EarliestFullBloom(plantTime, growTime)

	elapsed := time.Since(start)
	fmt.Printf("TestEarliestFullBloom5 Execution time: %s\n", elapsed)

	// Check the result
	if result != expected {
		t.Errorf("For plantTime %v and growTime %v, expected %d but got %d",
			plantTime, growTime, expected, result)
	}
}
