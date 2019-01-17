package main

import (
	"flag"
	"log"
	"packmule/data"
)

func main() {
	exportToYaml := flag.Bool("e", false, "Create Sample Yaml")
	testYaml := flag.Bool("t", false, "Test Yaml File")
	copyToDirectory := flag.Bool("c",false,"Copy to Destination Folder")
	exportToFileName := flag.String("f", "", "Yaml File Name")
	srcFolder := flag.String("sc","", "Source Folder for Copy")
	flag.Parse()

	if *testYaml == true && *exportToFileName != "" {
		data.ImportYamlFile(*exportToFileName)
	}

	if *exportToYaml == true && *exportToFileName != "" {
		data.GenerateExampleYaml(*exportToFileName)
	}

	if *copyToDirectory == true && *srcFolder != "" {
		d := data.ImportYamlFile(*exportToFileName)
		log.Printf("packmule: Copying From: %s To: %s",*srcFolder,d.HomeFolder)
		data.CopyLocalFiles(d, *srcFolder)
	}
}
