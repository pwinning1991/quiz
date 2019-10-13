package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func openFile(name string) *os.File {
	fmt.Println("calling the openFile Func")
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("unable to open CSV file %s", name)
	}
	return file

}

func readCsv(file io.Reader) [][]string {
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	return lines

}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv in the format of 'question,answer'")

	flag.Parse()
	file := openFile(*csvFilename)

	lines := readCsv(file)

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
