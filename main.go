package main

import (
	"flag"
	"log"
	"os"

	"github.com/mister-turtle/nmap-parser/nmap"
)

func main() {

	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	argXMLFile := flag.String("x", "", "NMap XML file to parse")
	argOutDir := flag.String("o", "./nmap-parser", "Output directory to use")
	flag.Parse()

	if *argXMLFile == "" {
		flag.Usage()
		log.Fatal("Please supply an XML file to parse")
	}

	if _, err := os.Stat(*argOutDir); os.IsNotExist(err) {
		err := os.MkdirAll(*argOutDir, 0777)
		if err != nil {
			log.Fatalf("couldnt create output dir: %v", err)
		}
	}

	outputter, err := nmap.NewOutputter(*argXMLFile, *argOutDir)
	if err != nil {
		log.Fatal(err)
	}

	err = outputter.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Output %s to %s\n", *argXMLFile, *argOutDir)
}
