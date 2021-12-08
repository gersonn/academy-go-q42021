package routes

import (
	"gobootcamp/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests(p controllers.PokemonController) *gin.Engine {
	r := gin.Default()

	r.GET("/_health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "app is working!",
		})
	})

	r.GET("/pokemon/:id", p.GetPokemonById)
	r.GET("/pokeapi", p.GetPokemonsFromPokeApi)
	r.POST("/pokemon/csv", p.ReadCsv)
	r.POST("/pokemon/workerpool", p.GetPokemonsWithWorkerPool)

	return r
	//r.Run() // listen and serve on 0.0.0.0:8080
}
