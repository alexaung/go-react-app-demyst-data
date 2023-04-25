//backend/middleware/corsMiddleware.go

package middleware

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true

    return cors.New(config)
}
