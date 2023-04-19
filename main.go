package main

import (
	"flag"
	"log"
)

var (
	file = flag.String("file", "", "file path")
)

func main() {
	flag.Parse()
	log.Println(file)
}
