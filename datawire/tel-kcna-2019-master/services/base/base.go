package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		preValue := now.Unix() % 10000
		if preValue < 1000 {
			// Get a 4-digit number
			preValue = 9999 - preValue
		}
		value := float64(preValue) / 10
		fmt.Fprint(w, strconv.FormatFloat(value, 'f', -1, 64))
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
