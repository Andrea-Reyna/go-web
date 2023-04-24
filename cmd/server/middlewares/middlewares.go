package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidateToken(ctx *gin.Context) {
	header := ctx.GetHeader("token")
	if header != os.Getenv("TOKEN") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}
	ctx.Next()
}

func Logger(ctx *gin.Context) {
	startTime := time.Now()

	ctx.Next()

	size := ctx.Writer.Size()

	method := ctx.Request.Method
	url := ctx.Request.URL.String()

	log.Printf("%s %s %d %s", method, url, size, time.Since(startTime))
}
