package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()

		// records is a [][]string that holds all the values read from the csv
		records, err := csv.NewReader(file).ReadAll()

		// logging the requestURI
		fmt.Printf(r.RequestURI);
		var response string
		switch r.RequestURI{
		case "/echo":
			//call echo method
			response = echo(records)
			break
		case "/invert":
			//call invert method
			response = invert(records)
			break;
		case "/flatten":
			response = flatten(records)
			break
		case "/sum":
			response = sum(records)
		case "/multiply":
			response = multiply(records)
		default:
			err = &http.ProtocolError{"Invalid operation"}
			break

		}

		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}

		fmt.Fprintln(w, response)
	})

	http.ListenAndServe(":8080", nil)
}

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
	for i = 0; i < r; i++ {
		for j = 0; j < c; j++ {
			transpose[i][j] = records[j][i]
		}
	}
	//returns
	return printMatrix(transpose)
}

func flatten(records [][]string) string {
	var response string
	var strs []string
	for _, row := range records {
		s := strings.Join(row, ",")
		strs = append(strs, s)
		//response = fmt.Sprintf("%s%s", response, strings.Join(row, ","))
	}
	response = strings.Join(strs, ",")
	return response
}

func sum (records [][]string) string {
	var sum int64
	for _, row := range records {
		for _, col := range row {
			i, err := strconv.Atoi(col)
			if (err != nil) {
				err = &http.ProtocolError{"Invalid input. Matrix does not contain valid integers"}
				return err.Error()
			}
			n := int64(i)
			sum = sum + n
		}
	}
	return strconv.FormatInt(sum, 10)

}

func multiply (records [][]string) string {
	var product int64 = 1
	for _, row := range records {
		for _, col := range row {
			i,err := strconv.Atoi(col)
			if(err != nil){
				err = &http.ProtocolError{"Invalid input. Matrix does not contain valid integers"}
				return err.Error()
			}
			n := int64(i)
			product = product * n
		}
	}
	return strconv.FormatInt(product, 10)
}

func printMatrix(records [][]string) string {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}
