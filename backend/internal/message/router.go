package message

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	api := r.Group("/api") 
	{
		messageHandler := NewMessageHandler()
		api.GET("/v1/message", messageHandler.GetMessage)
	}

	return r
}
