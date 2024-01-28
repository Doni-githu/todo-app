package person

import (
	"net/http"

	"github.com/Doni-githu/todo-app/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) DeletePerson(ctx *gin.Context) {
	id := ctx.Param("id")

	var person models.Person

	if result := h.db.First(&person, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}	

	h.db.Delete(&person)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Person has been deleted",
	})
}
