package v1

import (
	"net/http"
	"strconv"
	"time"

	"test-task/common/http/response"
	"test-task/internal/models"
	"test-task/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler interface {
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetAllTasks(c *gin.Context)
	StartTask(c *gin.Context)
	StopTask(c *gin.Context)
	GetWorkloads(c *gin.Context)
}

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) TaskHandler {
	return &taskHandler{
		taskService: taskService,
	}
}

// CreateTask creates a new task based on data received from the request.
func (h *taskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	createdTask, err := h.taskService.CreateTask(&task)
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

// UpdateTask updates task data based on data received from the request.
func (h *taskHandler) UpdateTask(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	task.ID = uint(taskID)
	updatedTask, err := h.taskService.UpdateTask(&task)
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask deletes a task based on the ID received from the request.
func (h *taskHandler) DeleteTask(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("id"))

	err := h.taskService.DeleteTaskByID(uint(taskID))
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}

// GetTaskByID retrieves a task based on the ID received from the request.
func (h *taskHandler) GetTaskByID(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("id"))

	task, err := h.taskService.GetTaskByID(uint(taskID))
	if err != nil {
		response.New(c).Error(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

// GetAllTasks retrieves a list of all tasks.
func (h *taskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// StartTask starts a task for a specific user based on user and task IDs received from the request.
func (h *taskHandler) StartTask(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	taskID, err := strconv.ParseUint(c.Param("task_id"), 10, 64)
	if err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	task, err := h.taskService.GetTaskByID(uint(taskID))
	if err != nil {
		response.New(c).Error(http.StatusNotFound, err)
		return
	}

	startTime := time.Now()

	err = h.taskService.StartTask(uint(userID), task.ID, startTime)
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task started"})
}

// StopTask stops a task for a specific user based on user and task IDs received from the request.
func (h *taskHandler) StopTask(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	taskID, err := strconv.ParseUint(c.Param("task_id"), 10, 64)
	if err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	task, err := h.taskService.GetTaskByID(uint(taskID))
	if err != nil {
		response.New(c).Error(http.StatusNotFound, err)
		return
	}

	endTime := time.Now()

	err = h.taskService.StopTask(uint(userID), task.ID, endTime)
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task stoped"})
}

// GetWorkloads retrieves workloads for a specific user within a specified date range.
func (h *taskHandler) GetWorkloads(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	if endDate.Before(startDate) {
		response.New(c).Error(http.StatusBadRequest, err)
		return
	}

	workloads, err := h.taskService.GetWorkloads(uint(userID), startDate, endDate)
	if err != nil {
		response.New(c).Error(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"workloads": workloads})
}
