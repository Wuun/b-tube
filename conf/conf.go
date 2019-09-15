package conf

import (
	"os"

	"github.com/joho/godotenv"
)

//Init initial the config of the server,etc mysql connection,redis connection.
func Init() {
	// get env from local page.
	godotenv.Load()
	//load yaml for error handle.
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	//connect to mysql
	Database(os.Getenv("MYSQL_DSN"))
	//connection to redis.
	Redis()
}
