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

	Local LocalStruct
	Mysql map[string]MysqlStruct
	Cos   CosStruct
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

	Local = C.Local
	Mysql = C.Mysql
	Cos = C.Cos

	// check if path ends with "/"
	if Cos.Path[len(Cos.Path)-1:] != "/" {
		Cos.Path += "/"
	}
	// not start with "/"
	if Cos.Path[0:1] == "/" {
		Cos.Path = Cos.Path[1:]
	}
	// check if path ends with "/"
	if Local.Dir[len(Local.Dir)-1:] != "/" {
		Local.Dir += "/"
	}

	// log.Println("Config loaded", C)
	log.Printf("Config loaded")
}
