package dbMain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"pizzeria-management-service/src/config"
	"pizzeria-management-service/src/tracer"
)

var Db *sql.DB

func ConnectToDb() bool {
	db, err := sql.Open("mysql", config.Settings.DbMain.Login+":"+config.Settings.DbMain.Password+"@/"+config.Settings.DbMain.DatabaseName)
	if err != nil {
		tracer.Error(err)
		return false
	}
	Db = db
	tracer.Debug("Connect to Db")
	return true
}
