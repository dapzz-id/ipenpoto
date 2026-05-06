package models

import (
	"time"
	"gorm.io/gorm"

	"ipenpoto/app/enums"
	"github.com/google/uuid"
)

type RegistrationRequest struct {
	ID         				uuid.UUID 				`gorm:"type:uuid;primaryKey"`
	UserID					uuid.UUID 				`gorm:"type:uuid;not null;index"`
	TargetRole				enums.UserRole 			`gorm:"type:varchar(20);not null;index"`
	Payload					string 					`gorm:"type:text"`
	DocumentURL 			string 					`gorm:"type:text"`
	Status					enums.ApprovalStatus 	`gorm:"type:varchar(20);not null;default:'PendingApprove'"`
	AdminNote				*string					`gorm:"type:varchar(255)"`
	ProcessedBy				*uuid.UUID   			`gorm:"type:uuid;index"`
	ProcessedAt				*time.Time
	CreatedAt 				time.Time 				`gorm:"autoCreateTime"`
	UpdatedAt 				time.Time 				`gorm:"autoUpdateTime"`
}

func (rr *RegistrationRequest) BeforeCreate(tx *gorm.DB) (err error) {
	if rr.ID == uuid.Nil {
		rr.ID = uuid.New()
	}
	return
}