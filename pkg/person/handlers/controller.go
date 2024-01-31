package person

import (
	"github.com/Doni-githu/todo-app/pkg/person/services"
	"github.com/gin-gonic/gin"
)

type handler struct {
	s *services.Service
}

func RegisterRoutes(r *gin.Engine, services *services.Service) {
	h := &handler{
		s: services,
	}

	routes := r.Group("/people")
	routes.POST("/", h.AddPerson)
	routes.GET("/", h.GetPeople)
	routes.GET("/{id}/", h.GetPerson)
	routes.PUT("/{id}/", h.UpdatePerson)
	routes.DELETE("/{id}/", h.DeletePerson)
}