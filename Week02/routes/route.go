package routes

import (
	"Week02/db"
	"Week02/handle"
	"github.com/gin-gonic/gin"
)

// Route service route
func Route(r *gin.Engine, client db.ServiceInterface) {
	handle := handle.NewHandleClient(client)
	r.GET("/api/v1/user/:Id", handle.GetUserNameByID)
}
