package main

import "testing"

func TestOne(t *testing.T) {
	one("test.txt")
}

func TestTwo(t *testing.T) {
	two("test.txt")
}

func TestIsSubDir(t *testing.T) {
	t.Log(isSubDir("/", "/a"))
}
