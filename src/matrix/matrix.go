package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)



func sum (records [][]string) string {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Calculating sum of elements of matrix")
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
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Calculating product of elements of matrix")
	if len(records) == 0 {
		return strconv.Itoa(0)
	}
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
