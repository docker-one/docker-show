package routers

import (
	"docker-show/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func InitRouter() (router *gin.Engine) {
	router = gin.Default() //gin.New() //
	//config := cors.DefaultConfig()
	//config.AllowAllOrigins = true
	//router.Use(cors.New(config))

	//跨域
	router.Use(Cors())
	//拦截所有请求 打印下
	//router.Use(middleware())

	//log -------------------------------------------
	// Disable Console Color, you don't need console color when writing the logs to file.
	//gin.DisableConsoleColor()
	//// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//-----------------------------------------------------------

	//渲染html页面
	// 静态资源加载，本例为css,js以及资源图片
	router.Static("/led", "views/static")
	router.StaticFS("/temp_data", http.Dir("./temp_data"))

	//fmt.Println("open", "http://localhost:8888/myproject/")
	//router.Static("/myproject", "view/")
	//router.Static("static/*")

	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	/* 获得授权Token */
	//router.GET("/v1/auth", api.GetAuth)

	t1 := controllers.DockerShow{map[string]string{}}
	fmt.Println(t1)
	//t2 := controllers.FileC{Mgo: db.InitMongoDB2()}
	//r1 := controllers.RoleC{Mgo: db.InitMongoDB2()}
	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})

	//use mongo db
	v1 := router.Group("/v1/docker/")
	{
		fmt.Println(v1)
		v1.POST("/imagesList", t1.GetDockerImagesList)
		v1.POST("/containersList", t1.GetContainersList)
		////v1.GET("/getRealTimeGas", t1.GetRealTimeGas)
		//v1.POST("/getRealTimeGas", t1.GetRealTimeGas)
		//v1.POST("/getCurrentData", t1.GetCurrentData)
		//v1.POST("/saveCurrentData", t1.SaveCurrentData)
		//v1.POST("/getTempData", t1.GetTempData)
		//v1.POST("/getConfigInfo", t1.GetConfigInfo)
		//v1.POST("/saveConfigInfo", t1.SaveConfigInfo)
		//v1.POST("/login", t1.Login)
		//v1.Use(jwt.JWT())
		//{
		//	//realtime data form sensor upload
		//	v1.GET("/getRealTimeGas", t1.GetRealTimeGas)
		//	//v1.POST("/getUserList", t1.GetUserList)
		//	//v1.POST("/update", t1.Update)
		//	//v1.POST("/delete", t1.DelUser)
		//	//v1.POST("/add", t1.AddUser)
		//}

	}

	//fileRouter := router.Group("/v1/file/")
	//{
	//	fileRouter.POST("upload", t2.UploadFileQiNiu)
	//}

	//cmd := exec.Command("explorer", "http://localhost:8888/myproject/")
	//err2 := cmd.Start()
	//if err2 != nil {
	//	fmt.Println(err2.Error())
	//}

	return
}
