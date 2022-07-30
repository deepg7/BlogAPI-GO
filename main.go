package main

import (
	"BlogAPI-GO/config"
	"BlogAPI-GO/controllers"

	"github.com/gin-gonic/gin"
)

// func init(){
// 	config.LoadEnvironmentVariables()
// 	config.ConnectToDB()
// }

func main() {
	config.LoadEnvironmentVariables()
	config.ConnectToDB()
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
	
    // c.JSON(200, gin.H{"data": "hello world"})    
  })

  r.GET("/users", func(c *gin.Context) {
	controllers.FindUsers(c)
    // c.JSON(200, gin.H{"data": "hello world"})    
  })

  r.POST("/user",func(c *gin.Context)  {
	controllers.CreateUser(c)
  })

  r.Run()
}