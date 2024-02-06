package main

import (
	"log"

	"github.com/Doni-githu/todo-app/pkg/common/db"
	"github.com/Doni-githu/todo-app/pkg/common/middlewares"
	person "github.com/Doni-githu/todo-app/pkg/person/handlers"
	"github.com/Doni-githu/todo-app/pkg/person/repository"
	"github.com/Doni-githu/todo-app/pkg/person/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_LOCAL_URL").(string)

	gin.SetMode("debug")
	r := gin.Default()
	db, err := db.Init(dbUrl)
	if err != nil {
		log.Fatalf(err.Error())
		return 
	}
	repo := repository.NewRepository(db)
	s := services.NewService(repo)
	person.RegisterRoutes(r, s)
	r.Use(middlewares.CORSMiddleware())
	
	r.Run(":" +port)
}
