package main

//go mod init api-service
import (
	"api-service/models"
	"fmt"
	"net/http"
	"time"

	"api-service/controllers"
	"github.com/gin-gonic/gin"
)

//TODO
// use https://github.com/go-saas/saas/blob/main/examples/gorm/main.go#L10
//Auth, RBAC
//multi-tenancy
//yaml parser for config
//logger for logging
//json logging
func main() {
	//r := gin.Default()
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	//}
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	//if production
	//export GIN_MODE=release
	//gin.SetMode(gin.ReleaseMode)
	models.ConnectDatabase() // new
	router.GET("/maaz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.PostUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.PATCH("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.Run()
}
