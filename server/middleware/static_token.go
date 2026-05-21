package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

// 没有显式配置 header-name 时使用的默认请求头字段
const defaultStaticTokenHeader = "token"

// StaticToken 通用静态 token 校验中间件。
// 请求 header 中必须携带 headerName 字段，且值等于 expected，否则直接拒绝。
// 适合服务间内部调用 / 简单的接口防护，不建议用于面向最终用户的接口。
func StaticToken(headerName, expected string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(headerName)
		if token == "" {
			response.FailWithMessage("缺少 "+headerName+" 请求头", c)
			c.Abort()
			return
		}
		// expected 为空视为未配置，一律拒绝，避免误把保护接口变成裸接口
		if expected == "" || token != expected {
			response.FailWithMessage(headerName+" 校验失败", c)
			c.Abort()
			return
		}
		c.Next()
	}
}

// StaticTokenAuth 读取 config.yaml 中 static-token 配置的便捷封装。
// 每次请求都会重新读取 GVA_CONFIG，配置热更新可即时生效。
func StaticTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := global.GVA_CONFIG.StaticToken
		header := cfg.HeaderName
		if header == "" {
			header = defaultStaticTokenHeader
		}
		StaticToken(header, cfg.Value)(c)
	}
}
