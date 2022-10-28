package middleware

import (
	"github.com/gin-gonic/gin"
	"my_go_project/model/common/response"
)

func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.FailWithMessage(errorToString(err), c)
				c.Abort()
			}
		}()
		c.Next()
	}
}
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
