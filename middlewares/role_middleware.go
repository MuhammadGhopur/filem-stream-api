package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware
// Cek apakah user memiliki role yang diizinkan
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// ambil role dari context (hasil JWT middleware)
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "role tidak ditemukan",
			})
			c.Abort()
			return
		}

		// cek apakah role user termasuk yang diizinkan
		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		// jika tidak cocok
		c.JSON(http.StatusForbidden, gin.H{
			"error": "akses ditolak",
		})
		c.Abort()
	}
}
