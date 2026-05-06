package models

import (
	"time"
	"gorm.io/gorm"
	
	"ipenpoto/app/enums"
	"github.com/google/uuid"
)

type CorporateProfile struct {
	ID         				uuid.UUID 				`gorm:"type:uuid;primaryKey"`
	OwnerID					uuid.UUID 				`gorm:"type:uuid;not null;unique;index"`
	CompanyName 			string 					`gorm:"not null;index"`
	CompanyAddress 			string 					`gorm:"type:varchar(255)"`
	CompanyPhone 			string 					`gorm:"type:varchar(20)"`
	NPWP 					string 					`gorm:"type:varchar(50)"`
	DocumentURL 			string 					`gorm:"type:text"`
	ApprovalStatus 			enums.ApprovalStatus 	`gorm:"type:varchar(20);default:'PendingApprove'"`
	ApprovedAt 				*time.Time
	ApprovedBy 				*uuid.UUID   			`gorm:"type:uuid;index"`
	RejectedReason 			*string					`gorm:"type:varchar(255)"`
	CreatedAt 				time.Time 				`gorm:"autoCreateTime"`
	UpdatedAt 				time.Time 				`gorm:"autoUpdateTime"`

	Owner      				User 					`gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Approver 				*User 					`gorm:"foreignKey:ApprovedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (cp *CorporateProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if cp.ID == uuid.Nil {
		cp.ID = uuid.New()
	}
	return
}