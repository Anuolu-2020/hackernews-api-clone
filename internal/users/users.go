package users

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/Anuolu-2020/hackernews-api-clone/internal/db"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Fatalf("Error while hashing password: %v", err)
	}

	newUser := db.Users{Username: user.Username, Password: hashedPassword}
	result := db.Db.Create(&newUser)
	if result.Error != nil {
		log.Fatalf("Error occurred while creating user: %v", result.Error.Error())
	}
}

func GetUserIdByUsername(username string) (int, error) {
	var result db.Users
	dbResult := db.Db.Model(db.Users{Username: username}).First(&result)

	if dbResult.Error != nil {
		log.Printf("Error while fetching user: %v", dbResult.Error.Error())
		return 0, dbResult.Error
	}

	return int(result.ID), nil
}

func (user *User) Authenticate() bool {
	var result db.Users
	dbResult := db.Db.Where(&db.Users{Username: user.Username}).First(&result)

	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			log.Printf("Login err: %v", dbResult.Error)
			return false
		} else {
			log.Printf("Error while fetching user: %v", dbResult.Error.Error())
		}
	}

	return CheckPasswordHash(user.Password, result.Password)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
