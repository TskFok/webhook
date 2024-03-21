package global

import (
	"gorm.io/gorm"
)

var DataBase *gorm.DB
var MysqlDsn string
var MysqlPrefix string
var Mode string
