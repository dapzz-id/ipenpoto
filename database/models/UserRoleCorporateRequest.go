package models

import (
	"time"
	"gorm.io/gorm"

	"ipenpoto/app/enums"
	"github.com/google/uuid"
)

type UserRoleCorporateRequest struct {
	ID         				uuid.UUID 				`gorm:"type:uuid;primaryKey"`
	UserID					uuid.UUID 				`gorm:"type:uuid;not null;index"`
	TargetRole				enums.UserRole 			`gorm:"type:varchar(20);not null;index"`
	Status 					enums.ApprovalStatus 	`gorm:"type:varchar(20);not null;default:'Pending'"`
	ApprovedBy 				*uuid.UUID   			`gorm:"type:uuid;index"`
	ApprovedAt 				*time.Time
	RejectedReason 			*string					`gorm:"type:varchar(255)"`
	CreatedAt 				time.Time 				`gorm:"autoCreateTime"`
	UpdatedAt 				time.Time 				`gorm:"autoUpdateTime"`

	User      				User 					`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Approver 				*User 					`gorm:"foreignKey:ApprovedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (urcr *UserRoleCorporateRequest) BeforeCreate(tx *gorm.DB) (err error) {
	if urcr.ID == uuid.Nil {
		urcr.ID = uuid.New()
	}
	return
}