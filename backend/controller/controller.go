package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(accounts gin.Accounts) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	router.Use(cors.New(config))
	router.Use(static.Serve("/", static.LocalFile("./public", false)))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.File("./public/index.html")
	})

	adminRouter := router.Group("/admin")
	{
		adminRouter.Use(gin.BasicAuth(accounts))

		adminRouter.POST("/login", login)
		adminRouter.POST("/entries", createEntry)
		adminRouter.GET("/entries", getAllEntries)
		adminRouter.PUT("/entries/:id", updateEntry)
		adminRouter.DELETE("/entries/:id", deleteEntry)
	}

	router.GET("/bingo/:userName$", bingoGame)

	return router
}
