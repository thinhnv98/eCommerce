package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Common struct {
}

func (_self Common) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"time":    time.Now(),
		"timeUTC": time.Now().UTC(),
	})
}
