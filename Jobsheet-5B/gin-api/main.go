package main

import (
	"net/http"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


/* Task */
type Product struct {
	ID string `json:"id"`
	Name string `json:"name" binding:"required,min=2"`
	Stock *int16 `json:"stock" binding:"required,gte=0"`
}

type ProductStore struct {
	mu sync.RWMutex
	products map[string]Product
} 

func NewProductStore() *ProductStore{
	return &ProductStore{
		products: make(map[string]Product),
	}
}

func (s *ProductStore) addProduct(c *gin.Context) {
	var req Product
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New()
	req.ID = id.String()

	s.mu.Lock()
	defer s.mu.Unlock()

	s.products[req.ID] = req

	c.JSON(http.StatusCreated, req)
}

func (s *ProductStore) getProduct(c *gin.Context) {
	id := c.Param("id")
	s.mu.RLock()
	defer s.mu.RUnlock()
	product, exists := s.products[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "data": product})
}

func (s *ProductStore) deleteProduct(c *gin.Context) {
	id := c.Param("id")
	s.mu.Lock()
	defer s.mu.Unlock()

	if _,exists := s.products[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	delete(s.products, id)
	c.JSON(http.StatusNoContent, gin.H{})
}
/* End Task */

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