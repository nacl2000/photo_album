package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
)

func CheckSeesion(c *gin.Context) error {
	session := sessions.Default(c)
	if session.Get("username") == nil {
		return fmt.Errorf("Didn't log in.")
    }
	return nil
}
