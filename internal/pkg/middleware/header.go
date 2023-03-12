package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NoCache 是一个 Gin 中间件，用来禁止客户端缓存 HTTP 请求的返回结果.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Cors 是一个 Gin 中间件，用来设置 options 请求的返回头，然后退出中间件链，并结束请求(浏览器跨域设置).
func Cors(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		// 必选，设置允许访问的域名
		c.Header("Access-Control-Allow-Origin", "*")
		// 必选，逗号分隔的字符串，表明服务器支持的所有跨域请求的方法
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		// 逗号分隔的字符串，表明服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段。如果浏览器请求包括 Access-Control-Request-Headers 字段，则此字段是必选的
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// Secure 是一个 Gin 中间件，用来添加一些安全和资源访问相关的 HTTP 头.
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}
