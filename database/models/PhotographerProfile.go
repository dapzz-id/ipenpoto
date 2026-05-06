package models

import (
	"time"
	"gorm.io/gorm"

	"ipenpoto/app/enums"
	"github.com/google/uuid"
)

type PhotographerProfile struct {
	ID		  			uuid.UUID   			`gorm:"type:uuid;primaryKey"`
	UserID				uuid.UUID   			`gorm:"type:uuid;not null;unique;index"`
	Bio					*string 				`gorm:"type:varchar(100)"`
	Specialization 		*string 				`gorm:"type:varchar(255);index"`
	BasePrice			*string	 				`gorm:"type:decimal(10,2);index"`
	TypePrice			*string 				`gorm:"type:varchar(50);index"`
	LocationCity		*string 				`gorm:"type:varchar(100);index"`
	LocationCountry		*string 				`gorm:"type:varchar(100);index"`
	LocationLatitude	*float64 				`gorm:"type:decimal(10,8)"`
	LocationLongitude	*float64 				`gorm:"type:decimal(11,8)"`
	Equipment			*string 				`gorm:"type:varchar(255)"`
	DocumentURL 		string 					`gorm:"type:text"`
	ApprovalStatus 		enums.ApprovalStatus 	`gorm:"type:varchar(20);default:'PendingApprove'"`
	ApprovedAt 			*time.Time 		
	ApprovedBy 			*uuid.UUID   			`gorm:"type:uuid;index"`
	RejectedReason 		*string					`gorm:"type:varchar(255)"`
	CreatedAt 			time.Time 				`gorm:"autoCreateTime"`
	UpdatedAt 			time.Time 				`gorm:"autoUpdateTime"`

	User     			User 					`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Approver 			*User 					`gorm:"foreignKey:ApprovedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (pp *PhotographerProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if pp.ID == uuid.Nil {
		pp.ID = uuid.New()
	}
	return
}