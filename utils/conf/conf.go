package conf

import (
	"bytes"
	_ "embed"
	"github.com/spf13/viper"
	"webhook/app/global"
)

//go:embed conf.yaml
var conf []byte

func InitConfig() {
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(bytes.NewReader(conf))

	if nil != err {
		panic(err)
	}

	global.MysqlDsn = viper.Get("mysql.dsn").(string)
	global.MysqlPrefix = viper.Get("mysql.prefix").(string)
	global.Mode = viper.Get("mode").(string)
}
