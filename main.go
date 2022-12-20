package main

import (
	"github.com/soxft/mysql-backuper/config"
	"github.com/soxft/mysql-backuper/core"
)

func main() {
	config.Init()
	core.Run()
}
