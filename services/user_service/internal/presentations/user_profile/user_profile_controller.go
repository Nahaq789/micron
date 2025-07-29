package userprofile

import (
	"context"
	"log/slog"
	"net/http"
	"user_service/internal/application/commands"
	userprofile "user_service/internal/application/use_cases/user_profile"

	"github.com/gin-gonic/gin"
)

type UserProfileController struct {
	logger      *slog.Logger
	editService *userprofile.EditProfile
}

func NewUserProfileController(l *slog.Logger, editService *userprofile.EditProfile) *UserProfileController {
	return &UserProfileController{
		logger:      l,
		editService: editService,
	}
}

func (e UserProfileController) EditUserProfile(c *gin.Context) {
	var command commands.EditProfileCommand
	if err := c.ShouldBind(&command); err != nil {
		e.logger.ErrorContext(c.Request.Context(), "ユーザプロフィール更新リクエストの値が不正です。", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx := context.Background()
	if err := e.editService.EditUserProfile(ctx, command); err != nil {
		e.logger.ErrorContext(c.Request.Context(), "ユーザプロフィールの更新に失敗しました。", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}
}
