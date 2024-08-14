package models

import (
	"encoding/json"
	"io/ioutil"

	"os"
)

type User struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var users []User

// LoadUsers จากไฟล์ JSON
func LoadUsers() error {
	file, err := os.Open("users.json")
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	if err := json.Unmarshal(byteValue, &users); err != nil {
		return err
	}

	return nil
}

// SaveUsers ไปยังไฟล์ JSON
func SaveUsers() error {
	file, err := os.Create("users.json")
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(byteValue); err != nil {
		return err
	}

	return nil
}

// GetUsers คืนค่าผู้ใช้ทั้งหมด
func GetUsers() []User {
	return users
}

// AddUser เพิ่มผู้ใช้ใหม่
func AddUser(user User) {
	users = append(users, user)
}
