package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTMiddleware
// Middleware ini bertugas:
// 1. Ambil token dari header Authorization
// 2. Validasi token JWT
// 3. Ambil user_id dan role dari token
// 4. Simpan ke context Gin
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// ===============================
		// 1️⃣ Ambil Authorization Header
		// ===============================
		authHeader := c.GetHeader("Authorization")

		// jika header kosong → user belum login
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header tidak ada",
			})
			c.Abort()
			return
		}

		// ===============================
		// 2️⃣ Validasi format Bearer
		// ===============================
		// format wajib: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "format authorization harus Bearer <token>",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// ===============================
		// 3️⃣ Parse & validasi JWT
		// ===============================
		secretKey := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

			// pastikan algoritma HMAC (HS256)
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}

			return secretKey, nil
		})

		// token invalid / expired
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token tidak valid atau expired",
			})
			c.Abort()
			return
		}

		// ===============================
		// 4️⃣ Ambil claims
		// ===============================
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "claims token tidak valid",
			})
			c.Abort()
			return
		}

		// ===============================
		// 5️⃣ Simpan ke Gin Context
		// ===============================
		// jwt menyimpan angka sebagai float64
		userID := uint(claims["user_id"].(float64))
		role := claims["role"].(string)

		c.Set("user_id", userID)
		c.Set("role", role)

		// lanjut ke controller
		c.Next()
	}
}
