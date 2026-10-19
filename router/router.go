package router

import (
	// "database/sql"
	// "net/http"

	// "github.com/gin-gonic/gin"

	// "ginframework/todo"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"ginframework/todo"
)

// ================ database/sql Method ================
// // ========= Router agar bisa edit Post, Get, dll di local host =========
// func SetupRouter(db *sql.DB) *gin.Engine {
// 	r := gin.Default()

// 	r.Use(func(c *gin.Context) {
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

// 		// Jawab preflight request OPTIONS dengan 200 OK
// 		if c.Request.Method == http.MethodOptions {
// 			c.AbortWithStatus(http.StatusOK)
// 			return
// 		}
// 		c.Next()
// 	})

// 	v1 := r.Group("/api/v1")
// 	todo.RegisterRoutes(v1, &todo.Handler{DB: db})

// 	return r
// }

// ================ GORM Method ================
// SetupRouter membuat router Gin lengkap dengan semua route.
// Dipakai oleh main.go dan oleh test.
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	todo.RegisterRoutes(v1, &todo.Handler{DB: db})

	return r
}
