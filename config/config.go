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
func Env_pro_load(){
  err := godotenv.Load("envfiles/production.env")
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
    Env_pro_load()
    CONNECT := os.Getenv("HEROKU_DB_URL")
    return DBMS, CONNECT
}
