package models

import "time"

type Profile struct {
  ID        int                 `json:"id" gorm:"primary_key:auto_increment"`
  Phone     string              `json:"phone" gorm:"type: varchar(255)"`
  Gender    string              `json:"gender" gorm:"type: varchar(255)"`
  Address   string              `json:"address" gorm:"type: text"`
  UserID    int                 `json:"user_id"`
  User      UserProfileResponse `json:"user"`
  CreatedAt time.Time           `json:"-"`
  UpdatedAt time.Time           `json:"-"`
}