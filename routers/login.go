package routers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
)

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
}

const userDatabase string = "/data/user/database.json"

func checkUserAccount(requester User) (bool, error) {
	var userDataList []User
	userData, err := ioutil.ReadFile(userDatabase)  
	if err != nil {  
		return false, err
	}

	if err = json.Unmarshal(userData, &userDataList); err != nil { 
		return false, err
	}
	for _,  user := range userDataList {  
        if requester == user {  
            return true, nil
        }  
    }
	return false, nil
}


func login(c *gin.Context) {
	var requester User
	if err := c.BindJSON(&requester); err != nil {
		c.String(404, "Login failed!")
		return
	}
	isInvalidUser, err := checkUserAccount(requester)
	if err != nil {
		c.String(500, "Failed to read user database.")
		return
	}
	if !isInvalidUser {
		c.String(404, "Invalid user account.")
		return
	}
	session := sessions.Default(c)
	session.Set("username", requester.Username)
	session.Save()
	c.String(200, fmt.Sprintf("Hello %s", session.Get("username")))
}

func AddLoginRoutes(router *gin.RouterGroup) {
	loginRouter := router.Group("/login")
	loginRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", "login")
	})
	loginRouter.POST("/api", login)
}
