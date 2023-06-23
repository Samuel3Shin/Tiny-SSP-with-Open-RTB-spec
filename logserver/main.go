package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const logFilePath = "./log.txt"

func LogHandler(w http.ResponseWriter, r *http.Request) {
	// Read the body of the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log the data
	logToFile(string(body))

	// Respond with a 200 OK
	fmt.Fprint(w, "OK")
}

func logToFile(logText string) {
	// Open the log file in append mode
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer f.Close()

	// Write to the log file
	_, err = io.WriteString(f, time.Now().Format(time.RFC3339)+": "+logText+"\n")
	if err != nil {
		log.Printf("Error writing to file: %v", err)
	}
}

func main() {
	http.HandleFunc("/", LogHandler)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
