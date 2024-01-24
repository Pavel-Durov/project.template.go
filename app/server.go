package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"p3ld3v.dev/template/app/domain"
)

var db = make(map[string]string)

func SetupRouter(dep *Dependencies) *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "I am alive!")
	})

	engine.GET("/user/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		user, err := dep.UserService.GetUser(idNum)
		if err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		}
	})
	engine.POST("/user", func(c *gin.Context) {
		var body domain.CreateUser
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON format"})
			return
		}
		user, err := dep.UserService.CreateUser(body.Name)
		if err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "big oops!"})
		}
	})

	return engine
}

func StartServer(dep *Dependencies) *http.Server {
	addr := fmt.Sprintf("%s:%s", dep.Config.Host, dep.Config.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: SetupRouter(dep),
	}
	return server
}
