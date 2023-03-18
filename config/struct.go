package config

type CStruct struct {
	Local ConfigStruct           `yaml:"Local"`
	Mysql map[string]MysqlStruct `yaml:"Mysql"`
	Cos   CosStruct              `yaml:"Cos"`
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

type CosStruct struct {
	Region string `yaml:"Region"`
	Bucket string `yaml:"Bucket"`
	Secret struct {
		Id  string `yaml:"ID"`
		Key string `yaml:"Key"`
	} `yaml:"Secret"`
	Path string `yaml:"Path"`
}
