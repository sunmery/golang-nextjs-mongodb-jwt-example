package main

import (
	"01/16/api"
	"01/16/database"
	"01/16/tools"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	server := gin.Default()
	server.Use(tools.Cors())
	database.Start()

	// 处理Mongodb异常
	if database.ConnErr != nil {
		panic(database.ConnErr)
	}

	// 关闭mongodb服务
	defer func() {
		if closeDBErr := database.Client.Disconnect(context.TODO()); closeDBErr != nil {
			panic(closeDBErr)
		}
	}()

	// ping DB
	if err := database.Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("mongodb server start success!")

	userGroup := server.Group("/user")
	{
		userGroup.POST("/auth", api.SetToken)

		userGroup.GET("/", api.GetToken)
	}

	ginErr := server.Run(":4000")
	if ginErr != nil {
		panic(ginErr)
	}
}
