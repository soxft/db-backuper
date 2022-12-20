package config

import (
	"flag"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	C          *CStruct
	configPath string
)

func Init() {
	flag.StringVar(&configPath, "c", "config.yaml", "specify config file path")
	flag.Parse()

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}

	cRaw, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}

	C = &CStruct{}
	err = yaml.Unmarshal(cRaw, C)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
	// log.Println("Config loaded", C)
	log.Printf("Config loaded")
}
