package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

// SetTraceId set trace_id
func SetTraceId(c *gin.Context) {
	traceId := uuid.NewV4().String()
	c.Set("trace_id", traceId)
	c.Writer.Header().Set("trace_id", traceId)
}

//AccessLog 自定义log
func AccessLog() gin.HandlerFunc {
	logClient := logrus.New()

	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println("err", err)
	}
	logClient.Out = src
	logClient.SetLevel(logrus.InfoLevel)
	apiLogPath := "api.log"
	logWriter, err := rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		traceId := uuid.NewV4().String()
		c.Set("trace_id", traceId)
		c.Writer.Header().Set("trace_id", traceId)

		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logClient.Infof("%s | %3d | %13v | %15s | %s  %s |",
			traceId,
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}
