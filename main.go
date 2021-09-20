package main

import (
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"github.com/gorilla/mux"
	"github.com/johan-lejdung/go-microservice-api-guide/internal/app"
	"github.com/johan-lejdung/go-microservice-api-guide/internal/db"
)

import "encoding/json"

type Result struct {
	SERVERNAME string `json:"serverName"`
    USER string `json:"user"`
    PASSWORD string `json: password""`
    DBNAME string `json:dbName`
}
var dbConfig Result

// 日志输出
func initLog() {
	file := "./logs/" +"error"+ ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[Error Log]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func initDB() {
	bytes, err := ioutil.ReadFile("./db.json")
	if err != nil {
		log.Println("readFile error", err)
	}
	err = json.Unmarshal([]byte(bytes), &dbConfig)
	if err != nil {
		log.Println("parse dbjson error", err)
	}
}

func main() {
	initLog()
	initDB()
	database, err := db.CreateDatabase(dbConfig.SERVERNAME, dbConfig.USER, dbConfig.PASSWORD, dbConfig.DBNAME)
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
