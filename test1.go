package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	engine := gin.Default()
	engine.LoadHTMLGlob("./html/*")
	engine.GET("/熙皇帝的登录界面", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		username:=c.PostForm("username")
		password:=c.PostForm("password")


		c.HTML(http.StatusOK,"register.html",gin.H{
			"username": username,
			"password": password,
		})
	})

	engine.Run()

}
