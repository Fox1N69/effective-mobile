package v1

import (
	"net/http"
	"strconv"
	"time"

	"test-task/common/http/response"
	"test-task/internal/models"
	"test-task/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

func (h *taskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := h.taskService.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

func (h *taskHandler) UpdateTask(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = uint(taskID)
	updatedTask, err := h.taskService.UpdateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (h *taskHandler) DeleteTask(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("id"))

	if err := h.taskService.DeleteTaskByID(uint(taskID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *taskHandler) GetTaskByID(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("id"))

	task, err := h.taskService.GetTaskByID(uint(taskID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *taskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

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

	c.Status(http.StatusOK)
}

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *taskHandler) GetWorkloads(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use YYYY-MM-DD"})
		return
	}

	if endDate.Before(startDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "End date cannot be before start date"})
		return
	}

	workloads, err := h.taskService.GetWorkloads(uint(userID), startDate, endDate)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch workloads"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"workloads": workloads})
}
