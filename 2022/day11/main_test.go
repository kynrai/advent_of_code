package main

import (
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	one("test.txt")
}

func TestTwo(t *testing.T) {
	two("test.txt")
}

func TestParseOp(t *testing.T) {
	parseOp("Operation: new = old * old")
}

func TestParseTest(t *testing.T) {
	parseTest("Test: divisible by 13")
}

func TestParseTrue(t *testing.T) {
	s := "If true: throw to monkey 2"
	t.Log(strings.Split(strings.TrimSpace(s), " ")[5])
}
