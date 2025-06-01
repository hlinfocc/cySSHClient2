package admin

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hlinfocc/cySSHClient2/assets"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/dbhandle"
	"github.com/hlinfocc/cySSHClient2/pkg/dao/entity"
	"github.com/hlinfocc/cySSHClient2/pkg/datavo"
	"github.com/hlinfocc/cySSHClient2/pkg/utils"
	jwtutils "github.com/hlinfocc/cySSHClient2/pkg/utils/jwtUtils"
)

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"msg"`
}

func checkPortStatus(port int) bool {
	// 监听 端口
	log.Println("check listener Port", port)
	listenerPort := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", listenerPort)
	if err != nil {
		// 如果监听失败，则说明端口已被占用
		return false
	}
	// 关闭监听器
	defer listener.Close()

	// 如果监听成功，则说明端口未被占用
	return true
}

func writePort(port int) {
	filePath := "/var/run/hlinfo-cyssh-server.port"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(port))
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}
}

/**
* 启动web服务
 */
func StartWebServer() {
	port := 31918
	for !checkPortStatus(port) {
		port = port + 1
	}
	writePort(port)
	log.Println("监听的端口:", port)
	httpPort := fmt.Sprintf(":%d", port)
	// 创建一个默认的Gin引擎
	router := gin.Default()

	// 使用嵌入的静态资源
	router.StaticFS("/static", assets.FileSystem)

	// 使用cookie存储session信息
	// store := cookie.NewStore([]byte("secret"))
	// router.Use(sessions.Sessions("PHPSESSIONID", store))

	// 定义一个GET请求的路由，根路径"/"
	router.GET("/", func(c *gin.Context) {
		// 在这里提供Vue.js的入口HTML文件
		indexHTML, err := assets.GetIndexHtml()
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read index.html: %v", err)
			return
		}
		c.Data(http.StatusOK, "text/html", indexHTML)
	})

	// 登录
	router.POST("/api/user/login", func(ctx *gin.Context) {
		var loginParams datavo.LoginParams
		if err := ctx.ShouldBindJSON(&loginParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(loginParams)
		if loginParams.UserName != "admin" || loginParams.Passwd != "123456" {
			res := Resp{
				Code: 402,
				Msg:  "用户名或密码错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}

		code, msg, user := dbhandle.UserLoginCheck(loginParams)
		tokenStr, e := jwtutils.GenJwtToken(user.Id, user.RealName, user.Account, user.Status, user.UserType, user.Role)
		if e != nil {
			res := Resp{
				Code: 500,
				Msg:  "生成token失败",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		token := datavo.TokenInfo{
			Token:    tokenStr,
			UserInfo: user,
		}
		res := datavo.UserResp[datavo.TokenInfo]{
			Code: code,
			Msg:  msg,
			Data: token,
		}
		ctx.JSON(http.StatusOK, res)
	})

	router.GET("/api/user/userInfo", func(ctx *gin.Context) {
		var hostParams datavo.HostParams
		token := ctx.GetHeader("Authorization")
		if token == "" {
			token = ctx.GetHeader("token")
		}
		log.Println(token)
		customClaims, e := jwtutils.JWTParse(token)
		if e != nil {
			ctx.JSON(http.StatusUnauthorized, Resp{
				Code: 401,
				Msg:  "鉴权失败",
			})
			return
		}

		log.Println(hostParams)
		user := entity.UserInfo{
			Id:       customClaims.UserID,
			Account:  customClaims.Account,
			Role:     customClaims.Role,
			RealName: customClaims.RealName,
			Status:   customClaims.Status,
			UserType: customClaims.UserType,
		}
		tokenData := datavo.TokenInfo{
			Token:    token,
			UserInfo: user,
		}
		ctx.JSON(http.StatusOK, tokenData)
	})

	router.POST("/api/hosts/list", func(ctx *gin.Context) {
		var hostParams datavo.HostParams
		if err := ctx.ShouldBindJSON(&hostParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(hostParams)

		ctx.JSON(http.StatusOK, dbhandle.QueryHostsList(hostParams.Page, hostParams.Limit, hostParams.Description, hostParams.HostIp, hostParams.HostExtent == 1))
	})

	router.POST("/api/hosts/insert", func(ctx *gin.Context) {
		var hostParams *entity.Sshhostlist
		if err := ctx.ShouldBindJSON(&hostParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(hostParams)
		ctx.JSON(http.StatusOK, dbhandle.Save(hostParams, true))
	})
	router.POST("/api/hosts/update", func(ctx *gin.Context) {
		var hostParams *entity.Sshhostlist
		if err := ctx.ShouldBindJSON(&hostParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(hostParams)
		ctx.JSON(http.StatusOK, dbhandle.Save(hostParams, false))
	})
	router.DELETE("/api/hosts/delete", func(ctx *gin.Context) {
		id := utils.String2Int(ctx.Query("id"))
		var res Resp
		if id == 0 {
			// 处理错误
			res = Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		rs := dbhandle.DeleteHostById(id)
		if rs {
			res = Resp{
				Code: 200,
				Msg:  "删除成功",
			}
		} else {
			res = Resp{
				Code: 500,
				Msg:  "删除失败",
			}
		}
		ctx.JSON(http.StatusOK, res)
	})

	router.POST("/api/keys/list", func(ctx *gin.Context) {
		var hostParams datavo.HostParams
		if err := ctx.ShouldBindJSON(&hostParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(hostParams)
		ctx.JSON(http.StatusOK, dbhandle.QueryKeysList(hostParams.Page, hostParams.Limit))
	})
	router.POST("/api/keys/insert", func(ctx *gin.Context) {
		var ckParams *entity.Sshkeylist
		if err := ctx.ShouldBindJSON(&ckParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(ckParams)
		ctx.JSON(http.StatusOK, dbhandle.SaveKeys(ckParams, true))
	})
	router.POST("/api/keys/create", func(ctx *gin.Context) {
		var ckParams *datavo.CreateSshKeyParams
		if err := ctx.ShouldBindJSON(&ckParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(ckParams)
		ctx.JSON(http.StatusOK, dbhandle.CreateKeys(ckParams))
	})
	router.DELETE("/api/keys/delete", func(ctx *gin.Context) {
		id := utils.String2Int(ctx.Query("id"))

		var res Resp
		if id == 0 {
			// 处理错误
			res = Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		rs, msg := dbhandle.DeleteKeyById(id, false)
		if rs {
			res = Resp{
				Code: 200,
				Msg:  msg,
			}
		} else {
			res = Resp{
				Code: 500,
				Msg:  msg,
			}
		}
		ctx.JSON(http.StatusOK, res)
	})
	router.GET("/api/home/count", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dbhandle.HomeCount())
	})
	router.POST("/api/hostExtent/list", func(ctx *gin.Context) {
		var hostParams datavo.HostParams
		if err := ctx.ShouldBindJSON(&hostParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(hostParams)

		ctx.JSON(http.StatusOK, dbhandle.QueryHostExtentList(hostParams.Page, hostParams.Limit))
	})
	router.POST("/api/hostExtent/insert", func(ctx *gin.Context) {
		var hostParams *entity.HostExtent
		if err := ctx.ShouldBindJSON(&hostParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(hostParams)
		ctx.JSON(http.StatusOK, dbhandle.SaveHostExtent(hostParams, true))
	})
	router.POST("/api/hostExtent/update", func(ctx *gin.Context) {
		var hostParams *entity.HostExtent
		if err := ctx.ShouldBindJSON(&hostParams); err != nil {
			// 处理错误
			res := Resp{
				Code: 404,
				Msg:  "参数错误",
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		log.Println(hostParams)
		ctx.JSON(http.StatusOK, dbhandle.SaveHostExtent(hostParams, false))
	})

	// 启动Gin服务器，监听端口
	router.Run(httpPort)

}
