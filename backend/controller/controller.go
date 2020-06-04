package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/koloo91/release-bingo/model"
	"github.com/koloo91/release-bingo/repository"
	"github.com/koloo91/release-bingo/service"
	"net/http"
)

func SetupRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	adminRouter := router.Group("/admin")
	{
		adminRouter.Use(gin.BasicAuth(gin.Accounts{
			"admin": "admin",
		}))

		adminRouter.POST("/login", login)
		adminRouter.POST("/entries", createEntry)
		adminRouter.GET("/entries", getAllEntries)
		adminRouter.PUT("/entries/:id", updateEntry)
		adminRouter.DELETE("/entries/:id", deleteEntry)
	}

	return router
}

func login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func createEntry(ctx *gin.Context) {
	var entryVo model.EntryVo
	if err := ctx.ShouldBindJSON(&entryVo); err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpError{Message: err.Error()})
		return
	}

	entry, err := service.CreateEntry(ctx.Request.Context(), model.EntryVoToEntity(&entryVo))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.HttpError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, model.EntryEntityToVo(entry))
}

func getAllEntries(ctx *gin.Context) {
	entries, err := service.GetEntries(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.HttpError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": model.EntryEntitiesToVos(entries)})
}

func updateEntry(ctx *gin.Context) {
	entryId := ctx.Param("id")

	var entryVo model.EntryVo
	if err := ctx.ShouldBindJSON(&entryVo); err != nil {
		ctx.JSON(http.StatusBadRequest, model.HttpError{Message: err.Error()})
		return
	}

	entry, err := service.UpdateEntry(ctx.Request.Context(), entryId, entryVo.Text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.HttpError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.EntryEntityToVo(entry))
}

func deleteEntry(ctx *gin.Context) {
	entryId := ctx.Param("id")
	if err := repository.DeleteEntry(ctx.Request.Context(), entryId); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.HttpError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, "")
}
