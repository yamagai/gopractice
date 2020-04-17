package config

import (
	"os"
  "log"

	"github.com/joho/godotenv"
)

func Env_dev_load() {
	err := godotenv.Load("envfiles/develop.env")
	if err != nil {
		log.Fatalln(err)
	}
}
func GetDBConfig() (string, string) {
	DBMS := "mysql"
	if os.Getenv("GO_ENV") == "develop" {
    Env_dev_load()
    USER := os.Getenv("MYSQL_USER")
		PASS := os.Getenv("MYSQL_PASSWORD")
		PROTOCOL := "(localhost)"
		DBNAME := os.Getenv("MYSQL_DATABASE")
		CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
		return DBMS, CONNECT
	}
    CONNECT := "b1a81319703cc0:0bfc2129@tcp(us-cdbr-iron-east-01.cleardb.net:3306)/heroku_929f712b4d9906e?parseTime=true"
    return DBMS, CONNECT
}
