package conf

import (
	"btube/cache"
	"os"

	"github.com/joho/godotenv"
)

//GlobalConf is a singleton of Conf
var GlobalConf *Conf

//Conf is config would ues in this project.
type Conf struct {
	DictionaryDRR  string
	MysqlDSN       string
	SessionSecrect string
	WebAddr        string
}

//Init initial the config of the server,etc mysql connection,redis connection.
func Init() {
	InitConf()
	//load yaml for error handle.
	if err := LoadLocales(GlobalConf.DictionaryDRR); err != nil {
		panic(err)
	}

	//connect to mysql
	Database(GlobalConf.MysqlDSN)
	cache.Redis()
}

//InitConf init the GlobalConf
func InitConf() {
	if err := godotenv.Load();err != nil {
		panic(err)
	}
	dictionaryAddr := os.Getenv("DICTIONARY_ADDR")
	mysqlDSN := os.Getenv("MYSQL_DSN")
	sessionSecrect := os.Getenv("SESSION_SECRET")
	webAddr := os.Getenv("WEB_ADDR")
	GlobalConf = &Conf{
		DictionaryDRR:  dictionaryAddr,
		MysqlDSN:       mysqlDSN,
		SessionSecrect: sessionSecrect,
		WebAddr:        webAddr,
	}
}
