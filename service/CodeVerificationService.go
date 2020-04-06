package service

import (
	"fmt"
	"os"
	"shuiOauth/config"
	"shuiOauth/model"
)

// GetCodeRequest 请求必须满足
type GetCodeRequest struct {
	ClientID     string `json:"client_id" form:"client_id" binding:"required"`
	RedirectURL  string `json:"redirect_uri" form:"redirect_uri" binding:"required"`
	ResponseType string `json:"response_type" form:"response_type" binding:"required"`
	State        string `json:"state" form:"state" binding:"required"`
}

// CodeVerificationService 查询数据库看看clienid是否正确
func (gcr *GetCodeRequest) CodeVerificationService() {
	var oauthclient model.OauthClient
	config.DB.Where("client_id = ?", gcr.ClientID).First(&oauthclient)
	if oauthclient.ClientID == "" {
		fmt.Println("该用户id不存在")
		os.Exit(2)
	}
	os.Exit(1)
}
