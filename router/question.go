package router

import (
	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/goquestionapi/api"
)

type QuestionRouter struct {
	api *api.QuestionAPI
}

func NewQuestionRouter(api *api.QuestionAPI) *QuestionRouter {
	return &QuestionRouter{api}
}

func (router *QuestionRouter) Setup(routerGroup *gin.RouterGroup) {
	group := routerGroup.Group("/questions")
	group.GET("", router.api.GetQuestions)
	group.POST("", router.api.HandleAnswers)
}
