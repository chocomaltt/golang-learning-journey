package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserReq struct {
	Name string `json:"name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
	Age int `json:"age" binding:"gte=0,lte=120"`
}

func createUser(c *gin.Context) {
	var req CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func listProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": []string{"coffee", "tea"}})
}

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.NewString()
		c.Set("reques_id", id)
		c.Header("X-Request-ID", id)
		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c* gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors.Last().Error()})
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(RequestID())
	r.Use(ErrorHandler())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		q := c.DefaultQuery("q", "")
		c.JSON(http.StatusOK, gin.H{"id": id, "q": q})
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/products", listProducts)
		v1.POST("/users", createUser)
	}

	r.Run(":8080")
}