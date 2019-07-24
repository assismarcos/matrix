package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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
		logger.Println("Request URI: ", r.RequestURI);
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

	//http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
	//
	//}

	http.ListenAndServe(":8080", nil)
}

