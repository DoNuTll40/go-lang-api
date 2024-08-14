package routes

import (
	"go-lang/api/auth"
	"go-lang/api/middleware"
	"go-lang/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// โหลดข้อมูลผู้ใช้จากไฟล์ JSONเมื่อเริ่มเซิร์ฟเวอร์
	err := models.LoadUsers()
	if err != nil {
		panic("Failed to load users: " + err.Error())
	}

	r.POST("/login", loginHandler)
	r.POST("/register", registerHandler)
	r.GET("/me", auth.AuthenticateToken, protectedHandler)
}

func loginHandler(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for _, user := range models.GetUsers() {
		if user.Username == input.Username && user.Password == input.Password {
			token, err := middleware.GenerateJWT(user.UserId, user.Username, user.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"token":   token,
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
}

func registerHandler(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if newUser.Username == "" || newUser.Password == "" || newUser.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please fill in all fields"})
		return
	}

	for _, user := range models.GetUsers() {
		if user.Username == newUser.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
	}

	// กำหนด userId ใหม่ให้กับผู้ใช้ใหม่
	newUser.UserId = getNextUserID()

	models.AddUser(newUser)

	response := gin.H{
		"userid":   newUser.UserId,
		"username": newUser.Username,
		"role":     newUser.Role,
	}

	err := models.SaveUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "result": response})

}

func getNextUserID() int {
	users := models.GetUsers()
	if len(users) == 0 {
		return 1
	}
	return users[len(users)-1].UserId + 1
}

// ฟังก์ชันจัดการเส้นทางที่ได้รับการป้องกัน
func protectedHandler(c *gin.Context) {
	userid := c.MustGet("userId").(int)
	username := c.MustGet("username").(string)
	role := c.MustGet("role").(string)
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the protected route!",
		"userid":  userid,
		"user":    username,
		"role":    role,
	})
}
