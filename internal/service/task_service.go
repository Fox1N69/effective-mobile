package service

import (
	"test-task/internal/models"
	"test-task/internal/repo"
	"time"
)

type TaskService interface {
	CreateTask(task *models.Task) (*models.Task, error)
	UpdateTask(task *models.Task) (*models.Task, error)
	DeleteTaskByID(id uint) error
	GetTaskByID(id uint) (*models.Task, error)
	GetAllTasks() ([]*models.Task, error)
	StartTask(userID, taskID uint, startTime time.Time) error
	StopTask(userID, taskID uint, endTime time.Time) error
	GetWorkloads(userID uint, startDate, endDate time.Time) ([]*models.Workload, error)
}

type taskService struct {
	taskRepo repo.TaskRepo
}

func NewTaskService(taskRepo repo.TaskRepo) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (s *taskService) CreateTask(task *models.Task) (*models.Task, error) {
	return s.taskRepo.Create(task)
}

func (s *taskService) UpdateTask(task *models.Task) (*models.Task, error) {
	return s.taskRepo.Update(task)
}

func (s *taskService) DeleteTaskByID(id uint) error {
	return s.taskRepo.DeleteByID(id)
}

func (s *taskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.taskRepo.FindByID(id)
}

func (s *taskService) GetAllTasks() ([]*models.Task, error) {
	return s.taskRepo.Tasks()
}

func (s *taskService) StartTask(userID, taskID uint, startTime time.Time) error {
	task, err := s.taskRepo.FindByID(taskID)
	if err != nil {
		return err
	}

	task.StartTime = startTime
	_, err = s.taskRepo.Update(task)
	if err != nil {
		return err
	}
	return err
}

func (s *taskService) StopTask(userID, taskID uint, endTime time.Time) error {
	task, err := s.taskRepo.FindByID(taskID)
	if err != nil {
		return err
	}

	task.EndTime = endTime

	duration := task.EndTime.Sub(task.StartTime)
	task.TotalHours = duration.Hours()

	_, err = s.taskRepo.Update(task)
	return err
}

func (s *taskService) GetWorkloads(userID uint, startDate, endDate time.Time) ([]*models.Workload, error) {
	return s.taskRepo.GetWorkloads(userID, startDate, endDate)
}
