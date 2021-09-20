package controllers

import (
	"database/sql"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)
var dbmap = initDb()
func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:Satvik26#@tcp(localhost:3306)/OlympicsFanzone")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
}
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// import(
// 	"database/sql"
// 	"sync"
// 	"time"

// 	"go.elastic.co/apm/module/apmsql"

// 	_ "go.elastic.co/apm/module/apmsql/mysql"
// 	"go.uber.org/zap"
// )

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
// func NewDBClient(config *DBConfig) *sql.DB {
// 	// url := config.DBUserName + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?multiStatements=true&parseTime=true"

// 	url := "root" + ":" + "Satvik26#" + "@tcp(" + "localhost" + ":" + "3306" + ")/" + "olympicFanzone" + "?multiStatements=true&parseTime=true"
// 	client, err := apmsql.Open("mysql", url)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// client.SetMaxIdleConns(config.DBMaxIdleConnections)
// 	// client.SetMaxOpenConns(config.DBMaxOpenConnections)
// 	// client.SetConnMaxLifetime(time.Minute * 10)
// 	return client
// }
