package logger

import (
	"bufio"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// LogFile holds the file
var LogFile *os.File

// Setup creates a multiwriter logger
func Setup() {

	var err error
	LogFile, err = os.OpenFile("data/logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	wrt := io.MultiWriter(os.Stdout, LogFile)
	log.SetOutput(wrt)

}

// CleanUp closes the file
func CleanUp() {

	// Close file
	LogFile.Close()

}

// Sync generates a file for logging
func Sync() {

	// Time stamp
	year, month, _ := time.Now().Date()
	str := strconv.Itoa(year) + " " + month.String()
	str = str + time.Now().Format("15:04:05 PM")
	str = "data/logs/" + str + ".gz"

	LogFile.Name()

	file, _ := os.Open(LogFile.Name())
	read := bufio.NewReader(file)
	data, _ := ioutil.ReadAll(read)

	compressedFile, _ := os.Create(str)

	w := gzip.NewWriter(compressedFile)

	w.Write(data)

	w.Close()
	file.Close()
	compressedFile.Close()
	os.Remove("data/logs")
	Setup()

}
