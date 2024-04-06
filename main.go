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

	// Show usage if flags are invalid
	if path == "" {
		flag.Usage()
		os.Exit(1)
	}

	// get our goal
	var response bool = validate.ResourceValidate(path)

	// print response
	fmt.Printf("%t\n", response)
}
