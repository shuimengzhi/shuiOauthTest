package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetCodeRequest 请求必须满足
type GetCodeRequest struct {
	ClientID     string `json:"client_id" form:"client_id" binding:"required"`
	RedirectURL  string `json:"redirect_uri" form:"redirect_uri" binding:"required"`
	ResponseType string `json:"response_type" form:"response_type" binding:"required"`
	State        string `json:"state" form:"state" binding:"required"`
}

// SetSession 把请求存到session里
func SetSession(c *gin.Context) {
	var gcr GetCodeRequest
	// 如果该填的四个参数都添了则执行括号内容
	if c.ShouldBind(&gcr) == nil {
		clientID := gcr.ClientID
		redirectURL := gcr.RedirectURL
		responseType := gcr.ResponseType
		state := gcr.State
		session := sessions.Default(c)
		session.Set("client_id", clientID)
		session.Set("redirect_uri", redirectURL)
		session.Set("response_type", responseType)
		session.Set("state", state)
		session.Save()

	}

}
