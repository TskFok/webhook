package router

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"webhook/app/global"
	"webhook/app/model"
)

var Handle *gin.Engine

type res struct {
	Data any    `json:"data,omitempty"`
	Type string `json:"type,omitempty"`
}

func InitRouter() {
	gin.SetMode(global.Mode)

	Handle = gin.New()
	Handle.Use(gin.Recovery())
	Handle.Use(gin.Logger())

	Handle.POST("/webhook", func(context *gin.Context) {
		r := &res{}
		decode := json.NewDecoder(context.Request.Body)
		decodeErr := decode.Decode(r)

		if decodeErr != nil {
			fmt.Println(decodeErr.Error())
		}

		log := model.NewLog()

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(r.Data)

		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		// 将buf转换为字符串
		jsonString := buf.String()

		log.Type = r.Type
		log.Data = jsonString
		id := log.Create(log)

		fmt.Println(id)
	})

}
