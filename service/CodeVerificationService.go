package service

import (
	"os"
	"shuiOauth/config"
	controller "shuiOauth/controller"
	"shuiOauth/model"

	"github.com/liudng/godump"
)

// CodeVerificationService 查询数据库看看clienid是否正确
func (gcr controller.GetCodeRequest) CodeVerificationService() {
	var oauthclient model.OauthClient
	a := config.DB.Where("client_id = ?", gcr.ClientID).First(&oauthclient)
	godump.Dump(a)
	godump.Dump(oauthclient)
	os.Exit(1)
}
