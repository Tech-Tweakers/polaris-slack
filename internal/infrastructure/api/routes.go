package api

import (
	"slack-messages-api/internal/domain/appcontext"
	"slack-messages-api/internal/infrastructure/slackclient"
	"slack-messages-api/internal/infrastructure/slackworker"

	"github.com/gin-gonic/gin"
)

// type Interface interface {
// 	Len() int
//	Less(i, j int) bool
//	Swap(i, j int)
//}

// type ByID []models.Messages

// func (a ByID) Len() int           { return len(a) }
// func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }
// func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func MakeDefaultRoutes(ctx appcontext.Context) {
	router := gin.Default()
	router.Use(setupCors())
	gin.SetMode("release")

	router.GET("/slack-messages", checkNewMessages)
	router.POST("/slack-reply", slackclient.GetPayloadFrontEnd)

	router.Run(":9990")
}

func checkNewMessages(c *gin.Context) {
	responsePayload := slackworker.CheckNewMessages()
	// sortedPayload := sortStruct(responsePayload)
	c.IndentedJSON(200, responsePayload)
}

// func sortStruct(responsePayload []models.Messages) []models.Messages {
// 	sort.SliceStable(responsePayload, func(i, j int) bool {
// 		return responsePayload[i].ID > responsePayload[j].ID
// 	})
// 	return responsePayload
// }

func setupCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
