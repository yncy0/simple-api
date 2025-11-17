package main

import (

	"github.com/gin-gonic/gin"
	"github.com/yncy0/simple-api/internal/message"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := message.SetupRouter()
	r.Run(":8080")
}
