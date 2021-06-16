package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string `form:"username" json:"user" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLGlob("templates/**/*")

	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON:返回JSON格式数据
		// c.JSON(200, gin.H{
		// 	"message": "Hello world!",
		// })
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})
	// RESTful
	r.GET("/book", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "GET",
		// })
		c.Request.URL.Path = "/book"
		c.Request.Method = "POST"
		r.HandleContext(c)
		// c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	//******************
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	// 使用自定义模板函数
	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
			"name":  "<h3>my name is : zhangsan<h3>",
			"path":  getCurrentPath(),
		})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		// 方法二：使用结构体传输JSON数据
		type message struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		msg := message{
			Name:    "小王子",
			Message: "Hello world!",
			Age:     18,
		}
		c.JSON(http.StatusOK, msg)
		// c.XML(http.StatusOK, msg)
	})
	// querystring指的是URL中?后面携带的参数，例如：/user/search?username=小王子&address=沙河。 获取请求的querystring参数的方法如下：
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	// 当前端请求的数据通过form表单提交时，例如向/user/search发送一个POST请求，获取请求数据的方式如下：
	r.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		// username := c.PostForm("username")
		// address := c.PostForm("address")
		// 使用参数绑定
		var user User
		if err := c.ShouldBind(&user); err == nil {
			fmt.Printf("user info:%#v\n", user)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		//输出json结果给调用方
		// c.JSON(http.StatusOK, gin.H{
		// 	"message":  "ok",
		// 	"username": username,
		// 	"address":  address,
		// })
	})

	r.GET("/file", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layouts/base.tmpl", gin.H{
			"time": time.Now(),
		})
	})
	// 文件上传
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		// file, err := c.FormFile("f1")
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }

		// log.Println(file.Filename)
		// dst := fmt.Sprintf("E:/ginDemo/ginDemo/tmp/%s", file.Filename)
		// // 上传文件到指定的目录
		// c.SaveUploadedFile(file, dst)
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		// })
		// 多文件上传
			// Multipart form
			form, _ := c.MultipartForm()
			files := form.File["file"]
	
			for index, file := range files {
				log.Println(file.Filename)
				dst := fmt.Sprintf("E:/ginDemo/ginDemo/tmp/%d_%s", index, file.Filename)
				// 上传文件到指定的目录
				c.SaveUploadedFile(file, dst)
			}
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("%d files uploaded!", len(files)),
			})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
func getCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "./"
}
