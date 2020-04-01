package model

import (
	"github.com/guregu/null"
)

type OauthClient struct {
	ClientID     string      `gorm:"column:client_id;primary_key" json:"client_id"`
	ClientSecret null.String `gorm:"column:client_secret" json:"client_secret"`
	RedirectURI  null.String `gorm:"column:redirect_uri" json:"redirect_uri"`
	GrantTypes   null.String `gorm:"column:grant_types" json:"grant_types"`
	Scope        null.String `gorm:"column:scope" json:"scope"`
	UserID       null.String `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (o *OauthClient) TableName() string {
	return "oauth_clients"
}
