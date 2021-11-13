package main

import (
	"github.com/gin-gonic/gin"
	"goServer/api"
	"goServer/global"
	"net/http"
)

func main(){
	global.InitViper() //初始化全局配置
	global.InitLog()	//初始化日志器
	//global.MyDB = global.InitDB()	//初始化数据库对象
	//fmt.Println(global.CF.Local.Qqdefaultpic)
	//fmt.Println(global.MyDB)

	startGin()
}

func startGin(){
	if !global.CF.Gin.Debug{
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.LoadHTMLGlob("../html/index.html")
	r.StaticFS("/static", http.Dir("../html/static"))
	r.StaticFS("/music", http.Dir("./music"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/migu/api", func(c *gin.Context) {
		kw := c.DefaultQuery("keyword", "抖音")
		c.JSON(200, api.MiguApi(kw))
	})

	r.GET("/netease/api", func(c *gin.Context) {
		kw := c.DefaultQuery("keyword", "抖音")
		if kw == "19723756" || kw == "3779629" || kw == "2884035" || kw == "3778678"{
			c.JSON(200, api.NeteaseApiTop(kw))
		}else{
			c.JSON(200, api.NeteaseApi(kw))
		}
	})

	r.GET("/qq/api", func(c *gin.Context) {
		kw := c.DefaultQuery("keyword", "抖音")
		c.JSON(200, api.QQApi(kw))
	})
	r.Run(global.CF.Gin.Host + ":" + global.CF.Gin.Port)
}
