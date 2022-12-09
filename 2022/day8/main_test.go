package main

import "testing"

func TestOne(t *testing.T) {
	one("test.txt")
}

func TestTwo(t *testing.T) {
	two("test.txt")
}

func TestVisible(t *testing.T) {
	t.Log(visible([]int{3, 0, 3, 7, 5}))
	t.Log(visible([]int{2, 5, 5, 1, 2}))
	t.Log(visible([]int{6, 5, 3, 3, 2}))
	t.Log(visible([]int{3, 3, 5, 4, 9}))
}
