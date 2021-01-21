package fun

import (
	"fmt"
	"math/rand"
	"testing"
)

func returnTwoInteger() (int, int) {
	return rand.Intn(5), rand.Intn(9)
}

func TestFunc(t *testing.T) {
	a, b := returnTwoInteger()
	fmt.Println(a, b)
}

func TestSlice(t *testing.T) {
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
}

func TestSlice2(t *testing.T) {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
}
