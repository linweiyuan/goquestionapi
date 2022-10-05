package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/goquestionapi/api"
	"github.com/linweiyuan/goquestionapi/middleware"
	"github.com/linweiyuan/goquestionapi/router"
	"github.com/linweiyuan/goquestionapi/sqlc"
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

	sqlDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("failed to connect to DB: [%v]", err)
	}

	db := sqlc.New(sqlDB)

	server = gin.Default()
	server.Use(middleware.CORS())

	rg := server.Group("/")
	router.NewQuestionRouter(api.NewQuestionAPI(db)).Setup(rg)
}

func main() {
	if err := server.Run(":" + config.ServerPort); err != nil {
		log.Fatalf("failed to start server: [%v]", err)
	}
}
