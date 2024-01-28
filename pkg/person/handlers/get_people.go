package person

import (
	"net/http"

	"github.com/Doni-githu/todo-app/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetPeople(ctx *gin.Context) {
	var people []models.Person

	if result := h.db.Find(&people).Limit(10); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, people)
}