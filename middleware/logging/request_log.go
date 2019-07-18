package logging

import (
	"bytes"
	"github.com/Han-Ya-Jun/gin_scaffold/pkg/logging"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

// 请求进入日志
func RequestInLog(c *gin.Context) {
	traceContext := lib.NewTrace()
	if traceId := c.Request.Header.Get("com-header-rid"); traceId != "" {
		traceContext.TraceId = traceId
	}
	if spanId := c.Request.Header.Get("com-header-spanid"); spanId != "" {
		traceContext.SpanId = spanId
	}

	c.Set("startExecTime", time.Now())
	c.Set("trace", traceContext)

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back

	lib.Log.TagInfo(traceContext, "_com_request_in", map[string]interface{}{
		"uri":    c.Request.RequestURI,
		"method": c.Request.Method,
		"args":   c.Request.PostForm,
		"body":   string(bodyBytes),
		"from":   c.ClientIP(),
	})
}

// 请求输出日志
func RequestOutLog(c *gin.Context, blew *bodyLogWriter) {
	// after request
	endExecTime := time.Now()
	//response, _ := c.Get("response")
	st, _ := c.Get("startExecTime")
	startExecTime, _ := st.(time.Time)
	logging.ComLogNotice(c, "_com_request_out", map[string]interface{}{
		"uri":       c.Request.RequestURI,
		"method":    c.Request.Method,
		"from":      c.ClientIP(),
		"response":  blew.body.String(),
		"proc_time": endExecTime.Sub(startExecTime).Seconds(),
	})
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestInLog(c)
		b := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		defer RequestOutLog(c, b)
		c.Writer = b
		c.Next()
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
