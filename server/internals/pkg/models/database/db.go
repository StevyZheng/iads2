package database

import (
	"github.com/jinzhu/gorm"
	"iads/server/internals/pkg/models/database/mysql"
	"iads/server/internals/pkg/models/database/pgsql"
	"iads/server/pkg/config"
)

var DBE *gorm.DB

func CreateDBEngine() {
	dbType := config.ConfValue.DBType
	if dbType == "pgsql" {
		DBE = pgsql.Eloquent
	} else if dbType == "mysql" {
		DBE = mysql.Eloquent
	}
}
