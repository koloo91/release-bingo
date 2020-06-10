package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/koloo91/release-bingo/service"
	"log"
	"net/http"
	"time"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func bingoGame(ctx *gin.Context) {
	userName := ctx.Param("userName$")
	wsHandler(ctx.Request.Context(), userName, ctx.Writer, ctx.Request)
}

func wsHandler(ctx context.Context, userName string, w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Failed to upgrade ws: %+v", err)
		return
	}

	go service.ConnectToGame(userName, conn)

	select {
	case <-ctx.Done():
		log.Println(conn.WriteControl(websocket.CloseMessage, []byte("server shutdown"), time.Now()))
	}
}
