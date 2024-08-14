package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func AuthenticateToken(c *gin.Context) {
	// ดึงค่า Authorization header จาก request
	jwtToken := c.GetHeader("Authorization")

	if !strings.HasPrefix(jwtToken, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization not bearer"})
		return
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(jwtToken, "Bearer "))

	// ตรวจสอบว่ามีค่า Authorization header หรือไม่
	if tokenString == "" {
		// ถ้าไม่มีค่า Authorization header หรือเป็นค่าเปล่า ส่ง response กลับไปว่า Unauthorized
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or malformed token"})
		c.Abort() // หยุดการทำงานของ request handler
		return
	}

	// แปลงค่า tokenString เป็น token และตรวจสอบความถูกต้อง
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil // ส่งคืน secret key สำหรับการตรวจสอบความถูกต้องของ token
	})

	// ตรวจสอบความถูกต้องของ token หรือเกิดข้อผิดพลาดในการแปลง token
	if err != nil || !token.Valid {
		// ถ้า token ไม่ถูกต้องหรือมีข้อผิดพลาด ส่ง response กลับไปว่า Unauthorized
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort() // หยุดการทำงานของ request handler
		return
	}

	// แปลง token.Claims เป็น type ของ Claims และตรวจสอบว่าเป็นไปตามที่คาดไว้หรือไม่
	claims, ok := token.Claims.(*Claims)

	fmt.Println(claims)
	if !ok {
		// ถ้า token claims ไม่ถูกต้อง ส่ง response กลับไปว่า Unauthorized
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort() // หยุดการทำงานของ request handler
		return
	}

	// ตั้งค่า username และ role ลงใน context ของ request เพื่อให้สามารถเข้าถึงได้ใน handler ถัดไป
	c.Set("userId", claims.UserID)
	c.Set("username", claims.Username)
	c.Set("role", claims.Role)

	// เรียกใช้ handler ถัดไปใน chain
	c.Next()
}
