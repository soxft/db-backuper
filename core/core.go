package core

import (
	"log"

	"github.com/soxft/mysql-backuper/backup"

	"github.com/soxft/mysql-backuper/config"
	// "github.com/soxft/mysql-backuper/tool"
)

func Run() {
	for k, v := range config.C.Mysql {
		if location, err := backup.Mysql(v.Host, v.Port, v.User, v.Pass, v.Db, config.C.Config.BackupDir); err != nil {
			log.Printf("%s > Backup error: %v", k, err)
		} else {
			log.Printf("%s > Backup created: %s", k, location)
		}
	}
}
