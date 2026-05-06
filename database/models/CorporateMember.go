package models

import (
	"time"
	"gorm.io/gorm"

	"ipenpoto/app/enums"
	"github.com/google/uuid"
)

type CorporateMember struct {
	ID         				uuid.UUID 						`gorm:"type:uuid;primaryKey"`
	CorporateID				uuid.UUID 						`gorm:"type:uuid;not null;index"`
	UserID					uuid.UUID 						`gorm:"type:uuid;not null;index"`
	RoleID					uuid.UUID 						`gorm:"type:uuid;not null;index"`
	Status 					enums.StatusUser			 	`gorm:"type:varchar(20);not null;default:'Pending'"`
	InvitedBy 				uuid.UUID 						`gorm:"type:uuid;index"`
	CreatedAt 				time.Time 						`gorm:"autoCreateTime"`
	UpdatedAt 				time.Time 						`gorm:"autoUpdateTime"`

	Corporate 				CorporateProfile 				`gorm:"foreignKey:CorporateID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      				User             				`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role      				UserRoleCorporateRequest    	`gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (cm *CorporateMember) BeforeCreate(tx *gorm.DB) (err error) {
	if cm.ID == uuid.Nil {
		cm.ID = uuid.New()
	}
	return
}