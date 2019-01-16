package main

import (
	"flag"
	"packmule/data"
)

func main() {
	exportToYaml := flag.Bool("e", false, "Create Sample Yaml")
	copyToDirectory := flag.Bool("c",false,"Copy to Destination Folder")
	exportToFileName := flag.String("f", "", "Yaml File Name")
	flag.Parse()



	if *exportToYaml == true {
		data.GenerateExampleYaml(*exportToFileName)
	}

	if *copyToDirectory == true {
		
	}
}
