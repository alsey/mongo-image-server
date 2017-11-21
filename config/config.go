package config

import (
	"github.com/alsey/mongo-image-server/logger"
	"os"
)

var (
	mongo_addr string
	mongo_db   string
	username   string
	password   string
	serv_port  string
)

func init() {

	serv_port = env("PORT0", "3000")

	mongo_addr = env("MONGO_ADDR", "127.0.0.1:27017")
	logger.Info("mongodb address is %s", mongo_addr)

	mongo_db = env("MONGO_DB", "img")
	logger.Info("mongodb database is %s", mongo_db)

	username = env("MONGO_USER", "")
	logger.Info("mongodb username is %s", username)

	password = env("MONGO_PASS", "")
	logger.Info("mongodb password is %s", password)

}

func GetMongoAddr() string {
	return mongo_addr
}

func GetMongoDB() string {
	return mongo_db
}

func GetServPort() string {
	return serv_port
}

func GetMongoUser() string {
	return username
}

func GetMongoPassword() string {
	return password
}

func env(nme string, def ...string) (val string) {
	val = os.Getenv(nme)
	if len(val) == 0 {
		if len(def) > 0 {
			val = def[0]
		} else {
			logger.Fatal("Missing environment variable " + nme + ".")
		}
	}
	return
}
