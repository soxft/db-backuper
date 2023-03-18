package main

import (
	"github.com/soxft/db-backuper/config"
	"github.com/soxft/db-backuper/core"
)

func main() {
	config.Init()
	core.Run()
}
