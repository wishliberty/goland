package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloword/webook/internal/repository"
	"helloword/webook/internal/repository/dao"
	"helloword/webook/internal/service"
	"helloword/webook/internal/web"
	"helloword/webook/internal/web/middleware"
	"strings"
	"time"
)

func main() {
	db := initDB()
	server := initWebServer() //解决跨域问题
	initUserHdl(db, server)
	server.Run(":8080")
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDao(db) //初始化
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	hdl := web.NewUserHandler(us)
	hdl.RegisterRoutes(server)
}
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
func initWebServer() *gin.Engine {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowCredentials: true,                                   //是否允许携带cookie等信息
		AllowHeaders:     []string{"Content-Type,Authorization"}, //必须加上Authorization
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	login := &middleware.LoginMiddleWareBuilder{}
	//session初始化
	store := cookie.NewStore([]byte("secret")) //存储数据的，也就是userid存哪里
	// 直接存在cookie
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())

	return server
}
