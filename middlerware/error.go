package middlerware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
)

func ErrorPrinter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if strerr, exists := c.Get("err"); exists {
			if err, ok := strerr.(error); ok {
				log.Printf("Type:[%T] Error:[%+v]\n", errors.Cause(err), err)
			}
		}
	}
}
