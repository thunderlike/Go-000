package handle

import (
	"log"
	"strconv"

	"Week02/db"
	"github.com/gin-gonic/gin"
)

// Handle Client server
type Handle struct {
	Client db.ServiceInterface
}

// NewHandleClient handle client
func NewHandleClient(client db.ServiceInterface) *Handle {
	return &Handle{Client: client}
}

// GetUserNameByID get user inforation
func (h *Handle) GetUserNameByID(c *gin.Context) {
	ID := c.Param("Id")
	id, err := strconv.Atoi(ID)
	if err != nil {
		log.Printf("get users info by userId strconv atoi id %v\n", err)
	}
	name, err := h.Client.GetUserNameByID(id)
	if err != nil {
		log.Printf("get users info by userId error %v\n", err)
	}
	c.JSON(200, gin.H{
		"data": name,
	})
}
