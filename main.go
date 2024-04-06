package main

import (
	"aiamrpv/validate"
	"flag"
	"fmt"
	"os"
)

func main(){
	// Define flags
	var path string
	
	// Parse flags
	flag.StringVar(&path, "path", "", "relative path to json file")

	flag.Parse()

	// Show usage if lags are invalid
	if path == "" {
		//fmt.Println("Please specify path to AWS::IAM::Role Policy json file")
		flag.Usage()
		os.Exit(1)
	}

	var response bool = validate.ResourceValidate(path)

	fmt.Printf("%t\n", response)
}