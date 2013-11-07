/**
 * User: jackong
 * Date: 11/5/13
 * Time: 4:50 PM
 */
package global

import (
	"github.com/gorilla/mux"
	"github.com/Jackong/log"
	"os"
	"fmt"
	"morning-dairy/config"
	"time"
	"github.com/gosexy/db"
	_ "github.com/gosexy/db/mysql"
	"github.com/Jackong/log/writer"
)

var (
	GoPath string
	Now func() string
	Time func() time.Time
	Project config.Config
	Log log.Logger
	Conn db.Database
	Router *mux.Router
)

func init() {
	baseEnv()
	loadConfig()
	openDb()

	Router = mux.NewRouter()

	fileLog := fileLog("access.log")
	mailLog := mailLog()
	Log = log.MultiLogger(fileLog, mailLog)
}

func baseEnv() {
	Time = func() time.Time {
		return time.Now()
	}

	Now = func() string {
		return Time().Format("2006-01-02 15:04:05")
	}

	GoPath = os.Getenv("GOPATH")
}

func loadConfig() {
	Project = config.NewConfig(GoPath  + "/src/morning-dairy/config/project.json")
}

func fileLog(name string) log.Logger {
	logFile, err := os.OpenFile(Project.String("log", "dir") + "/" + name, os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return log.NewLogger(logFile, Project.String("server", "name"), int(Project.Get("log", "file", "level").(float64)))
}

func mailLog() log.Logger {
	mailLog := &writer.Email{
		User: Project.String("log", "email", "user"),
		Password: Project.String("log", "email", "password"),
		Host : Project.String("log", "email", "host"),
		To : Project.String("log", "email", "to"),
		Subject: Project.String("log", "email", "subject"),
	}
	return log.NewLogger(mailLog, Project.String("server", "name"), int(Project.Get("log", "email", "level").(float64)))
}

func openDb() {
	settings := db.DataSource{
		Socket:   Project.String("db", "socket"),
		Database: Project.String("db", "database"),
		User:     Project.String("db", "user"),
		Password: Project.String("db", "password"),
	}

	var err error
	if Conn, err = db.Open("mysql", settings); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
