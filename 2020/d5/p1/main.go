package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func seatId(s string) int {

	rowMin := 0
	rowMax := 127
	colMin := 0
	colMax := 7

	for _, c := range s {
		if c == 'F' {
			rowMax = rowMax - ((rowMax - rowMin + 1) / 2)
		}
		if c == 'B' {
			rowMin = rowMin + ((rowMax - rowMin + 1) / 2)
		}
		if c == 'L' {
			colMax = colMax - ((colMax - colMin + 1) / 2)
		}
		if c == 'R' {
			colMin = colMin + ((colMax - colMin + 1) / 2)
		}
		// fmt.Println(rowMin, rowMax, colMin, colMax)
	}

	return rowMin*8 + colMax
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f))
}

func processInput(f io.Reader) string {
	startTime := time.Now().Unix()
	result := 0
	sId := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		sId = seatId(s.Text())
		if sId > result {
			result = sId
		}

	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	endTime := time.Now().Unix()
	fmt.Printf("Calculated result %v in %d seconds\n", result, endTime-startTime)

	return fmt.Sprintf("%d", result)
}
