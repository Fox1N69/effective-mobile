package v1

import (
	"net/http"
	"strconv"
	"test-task/common/http/response"
	"test-task/internal/models"
	"test-task/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetAllUsers(c *gin.Context)
	UsersWithFiltersAndPagination(c *gin.Context)
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

func (h *userHandler) UsersWithFiltersAndPagination(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageSize, _ := strconv.Atoi(pageSizeStr)

	users, err := h.service.UsersWithFiltersAndPagination(models.UserFilters{}, models.Pagination{Page: page, PageSize: pageSize})
	if err != nil {
		response.New(c).Error(501, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
