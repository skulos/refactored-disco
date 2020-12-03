package main

import (
	"flag"
	"os"

	logger "github.com/sirupsen/logrus"
	"mn8.ee/dns/v1/src/dns"
)

//"mn8.ee/dns/v1/dns-utility"

func main() {

	// Flags
	debugFlag := flag.Bool("d", false, "Prints the debug logs")
	fileFlag := flag.String("f", "test.yaml", "Name of the YAML file")

	// Flag Processing
	flag.Parse()

	// Checking the debug status
	if *debugFlag == true {
		logger.SetLevel(logger.DebugLevel)
	}

	// https://golangcode.com/check-if-a-file-exists/
	if fileExists(*fileFlag) {
		logger.Println("Example file exists")
	} else {
		logger.Println("Example file does not exist (or is a directory)")
	}

	//file, _ := os.Create(*fileFlag)

	// ip := dns.IP4{172, 23}
	// var t dns.A = dns.A{"host", [4]byte{172, 23, 138, 254}}

	// fmt.Debug(t)
	// logger.Print(ip)
	//_, _ = file.WriteString("This is a test \n blah \n     blah")
	c := dns.Config{Name: "String"}
	logger.Debug(c)
	e := c.ImportConfigs()

	if e != nil {
		logger.Print("No Error")
	}

	logger.Debug(c)

}

// https://golangcode.com/check-if-a-file-exists/
// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	logger.Print(!info.IsDir())
	return !info.IsDir()
}
