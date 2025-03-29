package utils

import (
	"log"
	"os"
	"strconv"
)

// GetPort returns the port number from the command line arguments and returns the default port 10000 in not provided.
func GetPort() (string, error) {
	var (
		err  error
		port string
	)

	switch len(os.Args) {
	case 2:
		port = os.Args[1]

		portNum, err := strconv.Atoi(port)
		if err != nil || portNum < 1024 || portNum > 65535 {
			log.Printf("Failed validating port: %v. Switching to default port 10000\n", err)
			port = ":10000"
		} else {
			port = ":" + port
		}
	default:
		port = ":10000"
	}
	return port, err
}
