package main

import (
	"github.com/pius-microservices/piopos-database/cmd"
	"log"
	"os"
)

func main() {
	err := cmd.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}