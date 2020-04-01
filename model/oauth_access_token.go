package model

import (
	"time"

	"github.com/guregu/null"
)

type OauthAccessToken struct {
	AccessToken string      `gorm:"column:access_token;primary_key" json:"access_token"`
	ClientID    string      `gorm:"column:client_id" json:"client_id"`
	UserID      null.String `gorm:"column:user_id" json:"user_id"`
	Expires     time.Time   `gorm:"column:expires" json:"expires"`
	Scope       null.String `gorm:"column:scope" json:"scope"`
}

// TableName sets the insert table name for this struct type
func (o *OauthAccessToken) TableName() string {
	return "oauth_access_tokens"
}
