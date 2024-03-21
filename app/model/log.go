package model

import (
	"fmt"
	"webhook/app/global"
)

type log struct {
	BaseModel
	Id   uint32 `gorm:"column:id;type:INT(10) UNSIGNED;AUTO_INCREMENT;NOT NULL"`
	Type string `gorm:"column:type;type:VARCHAR(255);"`
	Data string `gorm:"column:data;type:LONGTEXT;"`
}

func NewLog() *log {
	return new(log)
}

func (*log) Create(is *log) uint32 {
	db := global.DataBase.Create(is)

	if db.Error != nil {
		fmt.Println(db.Error.Error())

		return 0
	}

	return is.Id
}
