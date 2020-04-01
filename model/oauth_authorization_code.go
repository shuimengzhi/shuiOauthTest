package model

import (
	"time"

	"github.com/guregu/null"
)

type OauthAuthorizationCode struct {
	AuthorizationCode string      `gorm:"column:authorization_code;primary_key" json:"authorization_code"`
	ClientID          string      `gorm:"column:client_id" json:"client_id"`
	UserID            null.String `gorm:"column:user_id" json:"user_id"`
	RedirectURI       null.String `gorm:"column:redirect_uri" json:"redirect_uri"`
	Expires           time.Time   `gorm:"column:expires" json:"expires"`
	Scope             null.String `gorm:"column:scope" json:"scope"`
	IDToken           null.String `gorm:"column:id_token" json:"id_token"`
}

// TableName sets the insert table name for this struct type
func (o *OauthAuthorizationCode) TableName() string {
	return "oauth_authorization_codes"
}
