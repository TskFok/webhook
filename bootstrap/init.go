package bootstrap

import (
	"webhook/app/global"
	"webhook/utils/conf"
	"webhook/utils/database"
)

func Init() {
	conf.InitConfig()

	global.DataBase = database.InitMysql()
}
