package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/natanaelrusli/library-api-gin/internal/constants"
)

const (
	MyDataKey = "my-data-key"
)

func AddContextMiddleware(c *gin.Context) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.Set(constants.MyDataKey, uuid)
	c.Next()
}
