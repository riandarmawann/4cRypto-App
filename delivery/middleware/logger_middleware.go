package middleware

import (
	"4crypto/model"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	// Open log file
	logFile, err := os.OpenFile("D:/BELVA/FinalProject4cRypto/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// create new logger instance
	logger := log.New(logFile, "", 0)

	return func(c *gin.Context) {
		// request start time
		startTime := time.Now()

		// process request
		c.Next()

		// request end time
		endTime := time.Now()

		// log request details
		logRequest := model.LogRequest{
			Latency:    endTime.Sub(startTime),
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			Message:    c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}

		logString := "\n" +
			"[GIN] " + endTime.Format("2006/01/02 - 15:04:05") + " " +
			strconv.Itoa(logRequest.StatusCode) + " " +
			logRequest.Latency.String() + " " +
			logRequest.ClientIP + " " +
			logRequest.Method + " " +
			logRequest.Path + " " +
			logRequest.Message
		logger.Println(logString)

		_, err := gin.DefaultWriter.Write([]byte(logString))

		if err != nil {
			fmt.Println(err)
		}
	}
}
