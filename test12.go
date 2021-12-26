package main
//ctrl b 看源码
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	engine := gin.Default()
	auth := func(c *gin.Context){
		//吃一个小饼干
		value, err := c.Cookie("gin_cookie")
		if err != nil {
			c.JSON(403,gin.H{
				"message":"认证失败！！！没有cookie咯",
			})
			c.Abort()//终止
		}else{
			//将获取到的cookie的值写入上下文
			c.Set("cookie",value)
			c.Next()//挂起继续向下走，然后执行完成下面的函数，会反过来最后执行该中间件
			v,_:=c.Get("next")
			fmt.Println(v)
		}
	}
	engine.POST("/login",func (c *gin.Context){
		username:=c.PostForm("username")
		password:=c.PostForm("password")
		if username == "xixizi" && password == "xiaocaiji"{
			c.SetCookie("gin_cookie", username, 3600, "/", "", false, true)
			c.JSON(200,gin.H{
				"msg": "当当当！！！login successfully！！！",
			})
		}else{
			c.JSON(403,gin.H{
				"message":"憨憨,账号密码输错啦！",
			})
		}
	})
	//在中间放入鉴权中间件
	engine.GET("/熙皇帝的登录界面",auth, func(c *gin.Context) {
		cookie,_:=c.Get("cookie")
		str:=cookie.(string)
		c.String(200,"哇哦，眼前的帅比居然能够登录熙皇帝的账号："+str)
		//测试next函数
		c.Set("next","test next")
	})

	engine.Run()

}

