package models

import (
	"time"
	"gorm.io/gorm"
	
	"ipenpoto/app/enums"
	"github.com/google/uuid"
)

type User struct {
	ID       			uuid.UUID 				`gorm:"type:uuid;primaryKey"`
	Name 				string 					`gorm:"not null;index"`
	Username 			string 					`gorm:"unique;not null;index"`
	Email 				string 					`gorm:"unique;not null;index"`
	Password 			string 					`gorm:"not null"`
	Role 				enums.UserRole 			`gorm:"type:varchar(20);not null;default:'Customer';index"`
	AvatarURL 			string 					`gorm:"type:varchar(255)"`
	EmailVerifiedAt 	*time.Time 				
	Status 				enums.StatusUser 		`gorm:"type:varchar(20);not null;default:'active'"`
	CreatedAt 			time.Time 				`gorm:"autoCreateTime"`
	UpdatedAt 			time.Time 				`gorm:"autoUpdateTime"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}