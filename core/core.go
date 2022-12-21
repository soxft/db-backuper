package core

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
	"github.com/soxft/db-backuper/backup"
	"github.com/soxft/db-backuper/config"
	"github.com/soxft/db-backuper/db"
	"github.com/soxft/db-backuper/tool"
)

func Run() {
	c := cron.New(cron.WithSeconds())

	for k, v := range config.Mysql {
		if _, err := c.AddFunc(v.Cron, cronFunc(k, v)); err != nil {
			log.Fatalf("%s > Add Cron error: %v", k, err)
		} else {
			log.Printf("%s > Cron added: %s", k, v.Cron)
		}
	}

	c.Start()

	// wait for interrupt signal to gracefully shutdown the server with
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	c.Stop()
	log.Println("Bye! :)")
}

func cronFunc(k string, v config.MysqlStruct) func() {
	return func() {
		go run(k, v)
	}
}

// backup main func
func run(name string, info config.MysqlStruct) {
	defer func() {
		// recover
		if err := recover(); err != nil {
			log.Printf("%s > Backup error: %v", name, err)
		}
	}()

	if info.BackupTo == nil {
		log.Printf("%s > BackupTo is empty", name)
		return
	}

	if location, err := db.MysqlDump(info.Host, info.Port, info.User, info.Pass, info.Db, config.Local.Dir); err != nil {
		log.Printf("%s > Backup error: %v", name, err)
	} else {
		log.Printf("%s > Backup created: %s", name, location)

		if isMethodContains(info.BackupTo, "cos") {
			if rlocation, err := backup.ToCos(location, info.Db); err != nil {
				log.Printf("%s > cos upload error: %v", name, err)
			} else {
				log.Printf("%s > cos upload success: %s", name, rlocation)
			}
		}

		if !isMethodContains(info.BackupTo, "local") {
			_ = os.Remove(location)
			log.Printf("%s > local backup removed: %s", name, location)
		}

		// remove local backup files if max file num is set
		tool.DeleteLocal(config.Local.Dir, config.Local.MaxFileNum)
	}
}

// isMethodContains check if method is in list
func isMethodContains(list []string, method string) bool {
	for _, v := range list {
		if v == method {
			return true
		}
	}
	return false
}
