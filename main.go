package main

import (
	"log"
	"os"

	"github.com/soxft/db-backuper/config"
	"github.com/soxft/db-backuper/core"
)

func main() {
	log.SetOutput(os.Stdout)
	config.Init()
	core.Run()
}
