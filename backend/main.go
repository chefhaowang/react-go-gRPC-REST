package main

import (
    "fmt"
	// "github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
    "backend/db_connection"
	// "net/http"
	// "time"
)

var jwtKey = []byte("secret123")

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func main(){
    // dbconnection.InitMySQL()

    uri := "mongodb://localhost:27017"

	err := dbconnection.ConnectMongoDB(uri)
	if err != nil {
		fmt.Println(err)
		return
	}



    // r := gin.Default()

    // // CORS middleware
    // r.Use(func(c *gin.Context) {
    //     c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    //     c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    //     c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    //     if c.Request.Method == "OPTIONS" {
    //         c.AbortWithStatus(204)
    //         return
    //     }
    //     c.Next()
    // })

    // r.POST("/api/login", loginHandler)
    // r.GET("/api/admin/dashboard", AuthMiddleware("admin"), func(c *gin.Context) {
    //     c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin!"})
    // })
    // r.GET("/api/user/profile", AuthMiddleware("user"), func(c *gin.Context) {
    //     c.JSON(http.StatusOK, gin.H{"message": "Welcome User!"})
    // })

    // r.Run(":8000")


}

// func AuthMiddleware(requiredRole string) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         authHeader := c.GetHeader("Authorization")
//         if authHeader == "" {
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
//             return
//         }

//         // Remove "Bearer " prefix
//         tokenStr := authHeader[len("Bearer "):]

//         claims := jwt.MapClaims{}
//         token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
//             return jwtKey, nil
//         })

//         if err != nil || !token.Valid {
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
//             return
//         }

//         // Check role
//         userRole, ok := claims["role"].(string)
//         if !ok || userRole != requiredRole {
//             c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied for your role"})
//             return
//         }

//         // Pass user info to next handler
//         c.Set("username", claims["username"])
//         c.Set("role", userRole)

//         c.Next()
//     }
// }

// func loginHandler(c *gin.Context) {
//     var creds Credentials

//     if err := c.BindJSON(&creds); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//         return
//     }

//     var storedPassword, role string

//     err := authentication.DB.QueryRow("SELECT password, role FROM users WHERE username = ?", creds.Username).Scan(&storedPassword, &role)

//     if err != nil {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
//         return
//     }

//     if creds.Password != storedPassword {
//         c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
//         return
//     }

//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//         "username": creds.Username,
//         "role":     role,
//         "exp":      time.Now().Add(2 * time.Hour).Unix(),
//     })

//     tokenString, err := token.SignedString(jwtKey)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
//         return
//     }

//     c.JSON(http.StatusOK, gin.H{"token": tokenString})
// }
