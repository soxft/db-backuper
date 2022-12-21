package core

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
	"github.com/soxft/mysql-backuper/config"
	"github.com/soxft/mysql-backuper/db"
)

func Run() {

	c := cron.New()

	for k, v := range config.C.Mysql {
		if _, err := c.AddFunc(v.Cron, func() {
			log.Printf("%s > Cron triggered", k)
			// backup(k, v)
		}); err != nil {
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

func backup(name string, info config.MysqlStruct) {
	if location, err := db.MysqlDump(info.Host, info.Port, info.User, info.Pass, info.Db, config.C.Config.BackupDir); err != nil {
		log.SetOutput(os.Stderr)
		log.Printf("%s > Backup error: %v", name, err)
		log.SetOutput(os.Stdout)
	} else {
		log.Printf("%s > Backup created: %s", name, location)
	}
}
