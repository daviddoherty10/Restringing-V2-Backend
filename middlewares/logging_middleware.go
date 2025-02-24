package middlewares

import (
	"Restringing-V2/entity"
	"Restringing-V2/internal/database"
	"bytes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggingMiddleware(db database.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		// Capture request info
		clientIP := ctx.ClientIP()
		userAgent := ctx.GetHeader("User-Agent")
		referrer := ctx.GetHeader("Referer")
		method := ctx.Request.Method
		url := ctx.Request.URL.String()
		headers := ctx.Request.Header

		// Log incoming request
		log.Println("---- Incoming Request ----")
		log.Println("Client IP:", clientIP)
		log.Println("User-Agent:", userAgent)
		log.Println("Referrer:", referrer)
		log.Println("Method:", method)
		log.Println("URL:", url)
		log.Println("Headers:", headers)
		log.Println("-------------------------")

		// Save request data to DB
		var currentLog entity.LoggingMiddleware
		currentLog.ClientIP = clientIP
		currentLog.UserAgent = userAgent
		currentLog.Referer = referrer
		currentLog.RequestMethod = method
		currentLog.RequestURL = url

		// Wrap response writer to capture response body
		responseBody := &bytes.Buffer{}
		customWriter := &CustomResponseWriter{body: responseBody, ResponseWriter: ctx.Writer}
		ctx.Writer = customWriter

		// Process request
		ctx.Next()

		// Capture response metadata
		currentLog.StatusCode = ctx.Writer.Status()
		currentLog.ResponseBodyStr = responseBody.String()
		currentLog.Duration = time.Since(startTime)

		if err := db.CreateLog(currentLog); err != nil {
			log.Println("Failed to Insert Record to logging table", err.Error())
		}

		// Log response details
		log.Println("---- Response Details ----")
		log.Println("Status Code:", currentLog.StatusCode)
		log.Println("Response Body:", currentLog.ResponseBodyStr)
		log.Println("Duration:", currentLog.Duration)
		log.Println("-------------------------")

	}
}
