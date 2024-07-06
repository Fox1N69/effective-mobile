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
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

// GetAllUsers retrieves a list of all users.
func (h *userHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.Users()
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// UsersWithFiltersAndPagination retrieves users with filters and pagination parameters.
func (h *userHandler) UsersWithFiltersAndPagination(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageSize, _ := strconv.Atoi(pageSizeStr)

	users, err := h.service.UsersWithFiltersAndPagination(models.UserFilters{}, models.Pagination{Page: page, PageSize: pageSize})
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser creates a new user based on data received from the request.
func (h *userHandler) CreateUser(c *gin.Context) {
	var input struct {
		PassportNumber string `json:"passportNumber"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	createdUser, err := h.service.CreateUser(input.PassportNumber)
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "create user success", "user": createdUser})
}

// UpdateUser updates user data based on data received from the request.
func (h *userHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	err := h.service.UpdateUser(uint(id), &user)
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user based on the ID received from the request.
func (h *userHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteUser(uint(id))
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
