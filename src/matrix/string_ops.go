package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func echo(records [][]string) string {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Echoing matrix")

	return printMatrix(records)
}

func invert(records [][]string) string {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Inverting matrix")
	transpose := make([][]string, len(records[0]))

	var i, j, r, c int
	r = len(records[0])
	c = len(records)
	//initializing slice within slice
	for i := range transpose {
		transpose[i] = make([]string, r)
	}
	// Transposing values in records into transpose matrix
	for i = 0; i < c; i++ {
		for j = 0; j < r; j++ {
			transpose[i][j] = records[j][i]
		}
	}
	//returns
	return printMatrix(transpose)
}

func flatten(records [][]string) string {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Flattening matrix")
	var response string
	var strs []string
	for _, row := range records {
		s := strings.Join(row, ",")
		strs = append(strs, s)
	}
	response = strings.Join(strs, ",")
	return response
}



func printMatrix(records [][]string) string {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
