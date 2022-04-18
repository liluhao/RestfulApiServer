package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//强制浏览器不使用缓存
//NoCache 是一个中间件函数，它附加标头以防止客户端缓存 HTTP 响应。
//NOCache是请求头与响应头里的字段；通过指定指令来实现缓存机制。缓存指令是单向的，这意味着在请求中设置的指令，不一定被包含在响应中
func NoCache(c *gin.Context) {
	//no-cache:在发布缓存副本之前，强制要求缓存把请求提交给原始服务器进行验证(协商缓存验证)。
	//no-store ：缓存不应存储有关客户端请求或服务器响应的任何内容，即不使用任何缓存。
	//max-age=0：设置缓存存储的最大周期，超过这个时间缓存被认为过期(单位秒)
	//must-revalidate:一旦资源过期（比如已经超过），在成功向原始服务器验证之前(即你也可以把 revalidate 理解成“再次校验”的意思：再次校验看看缓存是不是真的过期了)，缓存不能用该资源响应后续请求
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	//一个 HTTP-日期 时间戳,即在此时候之后，响应过期;如果在Cache-Control响应头设置了 "max-age" 或者 "s-max-age" 指令，那么头会被忽略
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	//其中包含源头服务器认定的资源做出修改的日期及时间。 它通常被用作一个验证器来判断接收到的或者存储的资源是否彼此一致
	//比当前时间早8个小时，一个 HTTP-日期 时间戳；HTTP中的时间均用国际标准时间表示，从来不使用当地时间
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Options 是一个中间件函数
//浏览器跨域OPTIONS请求设置
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS") //该字段的值表明了服务器支持的所有 HTTP 方法：
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

//一些安全设置
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}

	// Also consider adding Content-Security-Policy headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}
