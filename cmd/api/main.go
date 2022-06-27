package main

import (
	"go-rest/internal/database"
	"go-rest/internal/exercise"
	"go-rest/internal/middleware"
	"go-rest/internal/user"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	db := database.NewDatabaseConn()
	exerciseService := exercise.NewExerciseService(db)
	userService := user.NewUserService(db)

	//exercise
	router.POST("/exercises", middleware.Authentication(userService), exerciseService.CreateExercise)
	router.GET("/exercises/:exerciseId", middleware.Authentication(userService), exerciseService.GetExercise)
	router.GET("/exercises/:exerciseId/score", middleware.Authentication(userService), exerciseService.GetUserScore)
	router.POST("/exercises/:exerciseId/questions", middleware.Authentication(userService), exerciseService.CreateQuestions)
	router.POST("/exercises/:exerciseId/questions/:questionId/answer", middleware.Authentication(userService), exerciseService.CreateAnswer)

	//user
	router.POST("/register", userService.Register)
	router.POST("/login", userService.Login)

	endless.ListenAndServe(":8000", router)
}
