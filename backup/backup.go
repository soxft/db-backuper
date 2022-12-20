package backup

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/soxft/mysql-backuper/tool"
)

func Mysql(host, port, user, password, databaseName, sqlPath string) (string, error) {
	// check if sqlPath dir exists
	if !tool.PathExists(sqlPath) {
		return "", errors.New("sqlPath does not exist")
	}

	var cmd *exec.Cmd

	backupPath := sqlPath + "db_" + databaseName + "_" + time.Now().Format("060102_150405") + ".sql"

	cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName, "--result-file="+backupPath)

	//stdout, _ := cmd.StdoutPipe()
	// defer stdout.Close()

	// stderr, _ := cmd.StderrPipe()
	// defer stderr.Close()

	if err := cmd.Start(); err != nil {
		log.Println(err)
		return "", err
	}

	// content, _ := ioutil.ReadAll(stderr)
	// log.Println(string(content))

	// wait for command to finish
	cmd.Wait()

	// check if the backup file is created or if file is 0 bytes
	if fi, err := os.Stat(backupPath); err == nil {
		if fi.Size() == 0 {
			log.Println("Backup file is 0 bytes")
			os.Remove(backupPath)
			return "", errors.New("backup error")
		}
	} else {
		log.Println(err)
		return "", err
	}

	return backupPath, nil
}
