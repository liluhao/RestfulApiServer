package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"regexp"
	"time"

	"apiserver/handler"
	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
	"github.com/zxmrlc/log"
)

/*为了演示中间件的功能，这里给apiserver新增功能：
日志记录每一个收到的请求*/
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		reg := regexp.MustCompile("(/v1/user|/login)") //这是标准库里的函数；func MustCompile(str string) *Regexp {}
		if !reg.MatchString(path) {                    //func (re *Regexp) MatchString(s string) bool {]
			return
		}
		//该中间件只记录业务请求，比如/v1/user和/login路径。
		// Skip for the health check requests.
		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		//该中间件需要截获HTTP的请求信息，然后打印请求信息，因为HTTP的请求Body,在读取过后会被置空，所以这里读取完后会重新赋值：
		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// The basic informations.
		method := c.Request.Method
		ip := c.ClientIP()

		//log.Debugf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		/*截获HTTP的Response更麻烦些，原理是重定向HTTP的Response到指定的IO流;
		截获HTTP的Request和Response后，就可以获取需要的信息，最终程序通过Iog.Infof()记录HTTP
		的请求信息。*/
		// get code and message
		var response handler.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			log.Errorf(err, "response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes())
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}
		//比如 2022-05-03 16:41:48.806 INFO middleware/logging.go:86 6.8713ms      | 127.0.0.1    | GET /v1/user | {code: 0, message: OK}
		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)
	}
}
