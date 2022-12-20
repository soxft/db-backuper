package config

type CStruct struct {
	Config ConfigStruct           `yaml:"Config"`
	Mysql  map[string]MysqlStruct `yaml:"Mysql"`
}

type ConfigStruct struct {
	BackupDir string   `yaml:"BackupDir"`
	BackupTo  []string `yaml:"BackupTo"`
}

type MysqlStruct struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
	User string `yaml:"User"`
	Pass string `yaml:"Pass"`
	Db   string `yaml:"Db"`
	Cron string `yaml:"Cron"`
}
