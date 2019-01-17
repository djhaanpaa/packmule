package main

import (
	"flag"
	"fmt"
	"log"
	"packmule/config"
	"packmule/data"
)

func main() {
	exportToYaml := flag.Bool("e", false, "Create Sample Yaml")
	testYaml := flag.Bool("t", false, "Test Yaml File")
	copyToDirectory := flag.Bool("c",false,"Copy to Destination Folder")
	serviceExport := flag.Bool("se",false,"Export Service File")
	yamlFileName := flag.String("f", "", "Yaml File Name")
	srcFolder := flag.String("sc","", "Source Folder for Copy")
	flag.Parse()

	config.GetIniFileName();

	if *testYaml == true && *yamlFileName != "" {
		d := data.ImportYamlFile(*yamlFileName)
		fmt.Printf("%+v\n\n",d)
	}

	if *exportToYaml == true && *yamlFileName != "" {
		data.GenerateExampleYaml(*yamlFileName)
	}

	if *serviceExport == true && *yamlFileName != "" {
		d := data.ImportYamlFile(*yamlFileName)
		data.ExportServiceFiles(d)
	}

	if *copyToDirectory == true && *srcFolder != "" {
		d := data.ImportYamlFile(*yamlFileName)
		log.Printf("packmule: Copying From: %s To: %s",*srcFolder,d.HomeFolder)
		data.CopyLocalFiles(d, *srcFolder)
	}
}
