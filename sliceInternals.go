package main

import "fmt"

func sliceMessage() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := []int{1, 2, 3, 4}

	fmt.Println(cap(s))
	fmt.Println(cap(s2))

	m := fmt.Sprintln("First:", s[0], s[0:1], s[:1])

	m += fmt.Sprintln("Last: ", s[len(s)-1], s[len(s)-1:], s[len(s)-1:])
	m += fmt.Sprintln("First 5", s[:5])

	m += fmt.Sprintln("Last 4:", s[5:])
	m += fmt.Sprintln("Middle 5:", s[2:7])
	return m
}

func getSlices() ([]int, []int, []int) {
	var s1 []int
	s2 := make([]int, 10)
	s3 := make([]int, 10, 50)

	return s1, s2, s3
}

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]

	s1[3] = 99

	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1

	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 4)
	copied := copy(s2, s1)
	fmt.Println(s2)
	s1[3] = 99
	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append(s1[0:0:0], s1...)
	s1[3] = 99
	return s1[3], s2[3]
}
