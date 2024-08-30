package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"
)

// Logger 使用Logger中间件时要关闭 gin.Recovery()
func Logger() gin.HandlerFunc {
	logPath := "log/log"

	err := os.MkdirAll(filepath.Dir(logPath), 0755)
	if err != nil {
		fmt.Println("err in make dir:", err)
	}

	logger := logrus.New()

	// 使用lumberjack进行日志切割
	logWriter := &lumberjack.Logger{
		Filename:   logPath + ".log", // 日志文件路径
		MaxSize:    128,              // 日志文件最大大小（MB）
		MaxBackups: 5,                // 保留旧文件的最大个数
		MaxAge:     30,               // 保留旧文件的最大天数
		Compress:   true,             // 是否压缩/归档旧文件
	}
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(logWriter)
	// 设置日志格式和时间戳格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	//writeMap := lfshook.WriterMap{
	//	logrus.InfoLevel:  logWriter,
	//	logrus.FatalLevel: logWriter,
	//	logrus.DebugLevel: logWriter,
	//	logrus.WarnLevel:  logWriter,
	//	logrus.ErrorLevel: logWriter,
	//	logrus.PanicLevel: logWriter,
	//}
	//Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})
	//
	//logger.AddHook(Hook)

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.WithFields(logrus.Fields{
					"error":       err,
					"stack_trace": fmt.Sprintf("%s", debug.Stack()),
				}).Error("Recovered from panic")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime).Milliseconds()
		spendTime := fmt.Sprintf("%d ms", stopTime)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"IP":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				entry.WithFields(logrus.Fields{
					"error": e.Error(),
				}).Error("请求处理出错")
			}
		}

		if statusCode >= 500 {
			entry.Error()
			for _, e := range c.Errors {
				entry.WithFields(logrus.Fields{
					"error": e.Error(),
				}).Error("请求处理出错")
			}
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
