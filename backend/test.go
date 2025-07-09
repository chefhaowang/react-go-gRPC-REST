package main

// import (
//     "github.com/gin-gonic/gin"
//     "net/http"
// )

// func main() {
//     r := gin.Default()

//     // CORS Middleware
//     r.Use(func(c * gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }
//         c.Next()
//     })

//     r.GET("/api/hello", func(c * gin.Context) {
//         c.JSON(http.StatusOK, gin.H{"message": "Hello from Go!"})
//     })

//     r.Run(":8000")
// }
