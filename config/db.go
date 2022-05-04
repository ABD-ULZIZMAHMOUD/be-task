package config

import (
	"be-task/helpers"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattn/go-sqlite3"
	"github.com/subosito/gotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"regexp"
	"strconv"
)

var DB *gorm.DB = nil
var err error = nil

func regex(re, s string) (bool, error) {
	return regexp.MatchString(re, s)
}

func ConnectToDatabase() {
	err = gotenv.Load()
	//sql register extended register regex function
	sql.Register("sqlite3_extended",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				return conn.RegisterFunc("regexp", regex, true)
			},
		})
	if helpers.CheckError(err) {
		panic(err)
	}
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_DATABASE"))
	log := logger.Default.LogMode(logger.Silent)
	if debug {
		log = logger.Default.LogMode(logger.Info)
	}

	f := sqlite.Dialector{
		DriverName: "sqlite3_extended",
		DSN:        os.Getenv("DATABASE_NAME"),
	}

	DB, err = gorm.Open(f, &gorm.Config{
		Logger: log,
	})
	if helpers.CheckError(err) {
		panic(err)
	}
}
