package db

import (
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/soxft/db-backuper/tool"
)

func MysqlDump(host, port, user, password, databaseName, sqlPath string) (string, error) {
	// check if sqlPath dir exists
	if !tool.PathExists(sqlPath) {
		return "", errors.New("sqlPath does not exist")
	}

	var cmd *exec.Cmd

	backupPath := sqlPath + "db_" + databaseName + "_" + time.Now().Format("060102_150405") + ".sql"

	cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName, "--result-file="+backupPath)

	//stdout, _ := cmd.StdoutPipe()
	// defer stdout.Close()

	stderr, _ := cmd.StderrPipe()
	defer stderr.Close()

	if err := cmd.Start(); err != nil {
		log.Println(err)
		return "", err
	}

	stderrContent, _ := io.ReadAll(stderr)

	// wait for command to finish
	cmd.Wait()

	// check if the backup file is created or if file is 0 bytes
	if fi, err := os.Stat(backupPath); err == nil {
		if fi.Size() == 0 {
			// log.Println("Backup file is 0 bytes")
			os.Remove(backupPath)
			return "", errors.New(string(stderrContent))
		}
	} else {
		return "", err
	}

	return backupPath, nil
}
