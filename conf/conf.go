package conf

import (
	"os"

	"github.com/joho/godotenv"
)

//GlobalConf is a singleton of Conf
var GlobalConf *Conf

//Conf is config would ues in this project.
type Conf struct {
	DictionaryDRR string
	MysqlDSN      string
	RedisDB       string
	RedisAddr     string
	RedisPW       string
}

//Init initial the config of the server,etc mysql connection,redis connection.
func init() {
	//load yaml for error handle.
	if err := LoadLocales(GlobalConf.DictionaryDRR); err != nil {
		panic(err)
	}

	//connect to mysql
	Database(GlobalConf.MysqlDSN)
	//connection to redis.
	//Redis()
}

//InitConf init the GlobalConf
func InitConf() {
	godotenv.Load()
	dictionaryAddr := os.Getenv("DICTIONARY_ADDR")
	mysqlDSN := os.Getenv("MYSQL_DSN")
	//redisDB := os.Getenv("REDIS_DB")
	//redisAddr := os.Getenv("REDIS_ADDR")
	//redisPW := os.Getenv("REDIS_PW")
	GlobalConf = &Conf{
		DictionaryDRR: dictionaryAddr,
		MysqlDSN:      mysqlDSN,
		//RedisDB:       redisDB,
		//RedisAddr:     redisAddr,
		//RedisPW:       redisPW,
	}
}
