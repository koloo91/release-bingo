package controller

import "github.com/gin-gonic/gin"

func SetupRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	adminRouter := router.Group("/admin")
	{
		adminRouter.Use(gin.BasicAuth(gin.Accounts{
			"admin": "admin",
		}))
		adminRouter.GET("/entries", getAllEntries)
		adminRouter.POST("/entries", createEntry)
		adminRouter.PUT("/entries/:id", updateEntry)
		adminRouter.DELETE("/entries/:id", deleteEntry)
	}

	return router
}

func getAllEntries(ctx *gin.Context) {

}

func createEntry(ctx *gin.Context) {

}

func updateEntry(ctx *gin.Context) {

}

func deleteEntry(ctx *gin.Context) {

}
