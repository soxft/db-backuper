package config

type CStruct struct {
	Local ConfigStruct           `yaml:"Local"`
	Mysql map[string]MysqlStruct `yaml:"Mysql"`
}

type ConfigStruct struct {
	Dir string `yaml:"Dir"`
}

type MysqlStruct struct {
	Host     string   `yaml:"Host"`
	Port     string   `yaml:"Port"`
	User     string   `yaml:"User"`
	Pass     string   `yaml:"Pass"`
	Db       string   `yaml:"Db"`
	Cron     string   `yaml:"Cron"`
	BackupTo []string `yaml:"BackupTo"`
}
