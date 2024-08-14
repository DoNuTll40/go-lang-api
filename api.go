package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var users = []User{
	{Username: "donut", Password: "donut1140", Role: "user"},
	{Username: "admin", Password: "admin", Role: "admin"},
	{Username: "superuser", Password: "superuser", Role: "superuser"},
}

func login() {
	r := gin.Default()

	r.Use(cors.Default())

	// login
	r.POST("/login", func(c *gin.Context) {
		var input User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		for _, user := range users {
			if user.Username == input.Username && user.Password == input.Password {
				c.JSON(http.StatusOK, gin.H{
					"message":  "Login successful",
					"role":     user.Role,
					"username": user.Username,
				})
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "ผู้ใช้งานหรือรหัสผ่านไม่ถูกต้อง"})
	})

	// register
	r.POST("/register", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if newUser.Username == "" || newUser.Password == "" || newUser.Role == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "โปรดป้อนข้อมูลให้ครบ"})
			return
		}

		for _, user := range users {
			if user.Username == newUser.Username {
				c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
				return
			}
		}

		users = append(users, newUser)
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

		fmt.Println(users)

	})

	r.Run(":8080") // Run on port 8080
}

func main() {
	login()
}
