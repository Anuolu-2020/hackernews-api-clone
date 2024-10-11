package links

import (
	"log"
	"strconv"
	"time"

	"github.com/Anuolu-2020/hackernews-api-clone/internal/db"
	"github.com/Anuolu-2020/hackernews-api-clone/internal/users"
)

type Link struct {
	ID        uint        `json:"id"`
	Title     string      `json:"title"`
	Address   string      `json:"address"`
	UserID    uint        `json:"user_id"`
	User      *users.User `json:"user"       gorm:"foreignKey:UserID"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (link Link) Save() int64 {
	userId, err := strconv.ParseUint(link.User.ID, 10, 64)
	if err != nil {
		log.Fatalf("Error occurred while converting string to uint: %v", err)
	}
	newLink := db.Links{
		Title:   link.Title,
		Address: link.Address,
		UserID:  uint(userId),
	}

	result := db.Db.Create(&newLink)
	if result.Error != nil {
		log.Printf("Failed to insert link: %v", result.Error.Error())
	}

	return int64(newLink.ID)
}

func GetAll() []Link {
	var links []Link

	err := db.Db.Model(&db.Links{}).Preload("User").Find(&links).Error
	if err != nil {
		log.Printf("Error while fetching links: %v", err)
	}

	return links
}
