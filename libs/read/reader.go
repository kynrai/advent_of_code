package read

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func InputAsStr(path string) []string {
	data := []string{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func InputAsInt(path string) []int {
	data := []int{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func InputAsFloat64(path string) []float64 {
	data := []float64{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		f, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, f)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func InputAsRunes(path string) [][]rune {
	data := [][]rune{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func InputAsBatches(path string) [][]string {
	data := [][]string{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	batcheStrs := strings.Split(string(b), "\n\n")

	for _, batch := range batcheStrs {
		lines := strings.Split(batch, "\n")
		data = append(data, lines)
	}

	return data
}

func InputAsChunks(path string, size int) [][]string {
	data := [][]string{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	chunk := []string{}
	batcheStrs := strings.Split(string(b), "\n")

	for i := 0; i < len(batcheStrs); i += 3 {
		chunk = append(chunk, batcheStrs[i], batcheStrs[i+1], batcheStrs[i+2])
		data = append(data, chunk)
		chunk = []string{}
	}

	return data
}

func InputAsSplitLines(path string, delim string) [][]string {
	data := [][]string{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), delim))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func Grid(path string) [][]int {
	data := [][]int{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ints := []int{}
		for _, v := range scanner.Text() {
			i, _ := strconv.Atoi(string(v))
			ints = append(ints, i)
		}
		data = append(data, ints)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}
