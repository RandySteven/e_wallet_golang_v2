package middleware

import (
	"assignment_4/logger"
	"bytes"
	"io"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		args := make(map[string]interface{})
		if c.Request.Body != nil {
			buf, _ := ioutil.ReadAll(c.Request.Body)
			rdr := ioutil.NopCloser(bytes.NewBuffer(buf))
			c.Request.Body = rdr
			args["request_body"] = c.Request.Body
		}

		logger.Log.WithFields(args).Info("Request")

		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		err := c.Errors.Last()
		if err == nil {
			logger.Log.Info("Success")
			return
		}

		args = make(map[string]interface{})
		args["method"] = c.Request.Method
		args["uri"] = c.Request.RequestURI
		args["status"] = c.Writer.Status()
		args["latency"] = endTime.Sub(startTime).Seconds()

		logger.Log.WithFields(args).Error(err)

	}
}
