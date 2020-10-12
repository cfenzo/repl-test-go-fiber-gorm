package dal

import (
  "time"
  "github.com/jinzhu/gorm"
  "github.com/gofrs/uuid"
)
// Base contains common columns for all tables.
type Base struct {
 ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
 CreatedAt time.Time
 UpdatedAt time.Time
 DeletedAt *time.Time `sql:"index"`
}
// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
 uuid, err := uuid.NewV4()
 base.ID = uuid
 return
}
