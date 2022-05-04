package providers

import (
	"be-task/config"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() {
	config.ConnectToDatabase()
}
