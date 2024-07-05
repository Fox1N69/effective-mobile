package v1

import (
	"test-task/common/http/response"
	"test-task/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetAllUsers(c *gin.Context)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{service: userService}
}

func (h *userHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.Users()
	if err != nil {
		response.New(c).Error(501, err)
		return
	}

	c.JSON(200, users)
}
