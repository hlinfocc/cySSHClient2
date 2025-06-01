package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hlinfocc/cySSHClient2/pkg/datavo"
	jwtutils "github.com/hlinfocc/cySSHClient2/pkg/utils/jwtUtils"
)

// JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			token = c.GetHeader("token")
		}
		log.Println(token)
		customClaims, e := jwtutils.JWTParse(token)
		if e != nil {
			c.JSON(http.StatusUnauthorized, datavo.SimpResp{
				Code: 401,
				Msg:  "鉴权失败",
			})
			return
		}
		c.Set("customClaims", customClaims)
		c.Next()
	}
}

// 中间件扩展：按域名后缀分组
func DomainGroupMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := strings.Split(c.Request.Host, ":")[0]
		// 根据域名处理相应内容
		c.Set("domain", domain)
		c.Next()
	}
}
