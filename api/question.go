package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/goquestionapi/sqlc"
	log "github.com/sirupsen/logrus"
)

type QuestionAPI struct {
	db *sqlc.Queries
}

func NewQuestionAPI(db *sqlc.Queries) *QuestionAPI {
	return &QuestionAPI{db}
}

func (api *QuestionAPI) GetQuestions(ctx *gin.Context) {
	questions, err := api.db.GetQuestions(ctx)
	if err != nil {
		log.Errorf("failed to get questions: [%v]", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"questions": questions,
	})
}
