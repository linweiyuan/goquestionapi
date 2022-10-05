package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/goquestionapi/middleware"
	"github.com/linweiyuan/goquestionapi/util"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

var (
	config util.Config
	server *gin.Engine
)

func init() {
	config = util.LoadConfig(".")

	server = gin.Default()
	server.Use(middleware.CORS())
}

func main() {
	if err := server.Run(":" + config.ServerPort); err != nil {
		log.Fatalf("failed to start server: [%v]", err)
	}
}
