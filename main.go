package main

import (
	"github.com/TanDat97/realestate-menu/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

var (
	engine *gin.Engine
	cfg    *viper.Viper
)

func init() {
	engine = gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/health"},
	}))

	cfg = config.GetConfig()
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
