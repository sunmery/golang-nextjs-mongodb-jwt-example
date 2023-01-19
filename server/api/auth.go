package api

import (
	"01/16/database"
	"01/16/tools"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Person struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type DatabasePerson struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

var JWTKEY = []byte("12312")
var tokenString string
var claim tools.Claims

func SetToken(c *gin.Context) {
	var person Person // 定义用户结构
	//处理转换JSON错误
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "请检查你的JSON结构!",
			"code": http.StatusBadRequest,
		})
		panic(err)
		return
	}

	var databasePerson DatabasePerson
	claim = tools.Claims{
		Username:       person.Username,
		StandardClaims: jwt.StandardClaims{},
	}
	var err3 error
	tokenString, err3 = tools.GenerateToken(claim, jwt.SigningMethodHS256, JWTKEY)
	if err3 != nil {
		panic(err3)
	}
	filter := bson.D{{"username", person.Username}}

	UserColl := database.Client.Database("golang_server").Collection("user")
	err2 := UserColl.FindOne(context.TODO(), filter).Decode(&databasePerson)
	if err2 != nil {
		panic(err2)
		return
	}

	if person.Username == databasePerson.Username {
		if person.Password == databasePerson.Password {
			c.JSON(http.StatusOK, gin.H{
				"msg":            "登录成功",
				"databasePerson": databasePerson,
				"person":         person,
				"tokenString":    tokenString,
				"code":           http.StatusOK,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "密码输入错误",
			"code": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  "账号或密码错误",
		"code": http.StatusBadRequest,
	})
}

func GetToken(c *gin.Context) {
	result, err := tools.ParseToken(tokenString, claim, JWTKEY)
	if err != nil {
		panic(err)
	}
	if result != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":         "登录成功",
			"tokenString": tokenString,
			"result":      result,
			"code":        http.StatusOK,
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "JWT校验错误",
		"code": http.StatusBadRequest,
	})
}
