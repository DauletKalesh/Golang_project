package main

import (
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("root:password@tcp(localhost:3306)/jwt_demo?parseTime=true") //root:dDGHHoRvoDIZaQTgr4my@containers-us-west-85.railway.app:7865/railway
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.POST("/movie/add", controllers.CreateMovie)
		api.DELETE("/movie/delete/:movieId", controllers.DeleteMovie)
		api.PUT("/movie/update/:movieId", controllers.UpdateMovie)
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
