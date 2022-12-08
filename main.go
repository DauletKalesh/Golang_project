package main

import (
	"fmt"
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	fmt.Println(URL)
	database.Connect(URL)
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}
//asem push
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/movie/add", controllers.CreateMovie)
		api.DELETE("/movie/delete/:movieId", controllers.DeleteMovie)
		api.PUT("/movie/update/:movieId", controllers.UpdateMovie)

		api.POST("/celebrity/add", controllers.CreateCelebrity)
		api.DELETE("/celebrity/delete/:celebrityId", controllers.DeleteCelebrity)

		api.PUT("/celebrity/update/:celebrityId", controllers.UpdateCelebrity)

		api.GET("/celebrity/get/:celebrityId", controllers.GetCelebrity)
		api.GET("/celebrities", controllers.GetCelebrites)

		api.GET("/movie/get/:movieId", controllers.GetMovie)
		// crud_movie := api.Group(("/movie")).Use()
		// {
		// 	crud_movie.POST("/movie/add", controllers.CreateMovie)
		// 	crud_movie.DELETE("/movie/delete/:movieId", controllers.DeleteMovie)
		// 	crud_movie.PUT("/movie/update/:movieId", controllers.UpdateMovie)
		// 	crud_movie.GET("/movie/get/:movieId", controllers.GetMovie)
		// }
		api.GET("/movies", controllers.GetMovies)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
