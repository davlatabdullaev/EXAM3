package api

import (
	"exam/api/handler"
	"exam/pkg/logger"
	"exam/service"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           BOOKS API
// @version         1.0
// @description     written api for uacademy exam.
func New(services service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.New(services, log)

	r := gin.New()

	r.Use(gin.Logger())
	{

		// // AUTHOR ENDPOINTS

		// r.POST("author", h.CreateAuthor)
		// r.GET("author/:id", h.GetAuthorByID)
		// r.GET("authors", h.GetAuthorList)
		// r.PUT("author/:id", h.UpdateAuthor)
		// r.DELETE("author/:id", h.DeleteAuthor)
		// r.PATCH("author/:id", h.UpdateAuthorPassword)

		// BOOK ENDPOINTS

		r.POST("/book", h.CreateBook)
		r.GET("/book/:id", h.GetBook)
		r.GET("/books", h.GetBookList)
		r.PUT("/book/:id", h.UpdateBook)
		r.DELETE("/book/:id", h.DeleteBook)
		r.PATCH("/book/:id", h.UpdateBookPageNumber)
		r.PATCH("/book_name/:id", h.UpdateBookName)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		r.Use(traceRequest)

	}

	return r
}

func traceRequest(c *gin.Context) {
	beforeRequest(c)

	c.Next()

	afterRequest(c)
}

func beforeRequest(c *gin.Context) {
	startTime := time.Now()

	c.Set("start_time", startTime)

	log.Println("start time:", startTime.Format("2006-01-02 15:04:05.0000"), "path:", c.Request.URL.Path)
}

func afterRequest(c *gin.Context) {
	
	startTime, exists := c.Get("start_time")
	if !exists {
		startTime = time.Now()
	}

	duration := time.Since(startTime.(time.Time)).Nanoseconds()

	log.Println("end time:", time.Now().Format("2006-01-02 15:04:05.0000"), "duration:", duration, "method:", c.Request.Method)
	fmt.Println()
}
