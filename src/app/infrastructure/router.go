package infrastructure

import (
	gin "github.com/gin-gonic/gin"

	"app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	blockchainController := controllers.NewBlockchainController()

	router.GET("/blockchain", func(c *gin.Context) { blockchainController.Show(c) })

	Router = router
}
