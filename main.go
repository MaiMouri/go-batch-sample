package main

import (
	"fmt"
	"github.com/MouriMai/go-batch-sample/batch/countryIp"
	"github.com/MouriMai/go-batch-sample/batch/hello"
	"github.com/MouriMai/go-batch-sample/batch/morning"
	"flag"
	// "net/http"
	// "regexp"
	// "strings"
)


func main() {
	fmt.Println("Go project started!")

	input := flag.String("s", "default", "input file name")
	// output := flag.String("output", "", "output file name")
	flag.Parse()

	if flag.Arg(0) == "countryIp" {
        countryIp.SubProcess()
    } else if flag.Arg(0) == "hello" {
        hello.Run(input)
    } else if flag.Arg(0) == "morning" {
        morning.Run(input)
    } 

}


