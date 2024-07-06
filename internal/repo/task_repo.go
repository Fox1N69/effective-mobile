package repo

import (
	"database/sql"
	"fmt"
	"test-task/internal/models"
	"time"

	"github.com/sirupsen/logrus"
)

type TaskRepo interface {
	Create(task *models.Task) (*models.Task, error)
	Update(task *models.Task) (*models.Task, error)
	DeleteByID(id uint) error
	FindByID(id uint) (*models.Task, error)
	Tasks() ([]*models.Task, error)
	GetWorkloads(userID uint, startDate, endDate time.Time) ([]*models.Workload, error)
}

type taskRepo struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewTaskRepo(db *sql.DB, log *logrus.Logger) TaskRepo {
	return &taskRepo{
		db:  db,
		log: log,
	}
}

func (r *taskRepo) Create(task *models.Task) (*models.Task, error) {
	const op = "repo.taskRepo.Create"
	query := `
		INSERT INTO tasks (user_id, name, description ,start_time, end_time, total_hours)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := r.db.QueryRow(query, task.UserID, task.Name, task.Description, task.StartTime, task.EndTime, task.TotalHours).Scan(&task.ID)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return task, nil
}

func (r *taskRepo) Update(task *models.Task) (*models.Task, error) {
	const op = "repo.taskRepo.Update"
	query := `
		UPDATE tasks
		SET user_id = $1, name = $2, start_time = $3, end_time = $4, total_hours = $5
		WHERE id = $6
	`

	_, err := r.db.Exec(query, task.UserID, task.Name, task.StartTime, task.EndTime, task.TotalHours, task.ID)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return task, nil
}

func (r *taskRepo) DeleteByID(id uint) error {
	const op = "repo.taskRepo.DeleteByID"
	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}

func (r *taskRepo) FindByID(id uint) (*models.Task, error) {
	const op = "repo.taskRepo.FindByID"

	query := `
		SELECT id, user_id, name, start_time, end_time, total_hours
		FROM tasks
		WHERE id = $1
	`

	var task models.Task
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.UserID, &task.Name, &task.StartTime, &task.EndTime, &task.TotalHours)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s %w", op, err)
		}
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return &task, nil
}

func (r *taskRepo) Tasks() ([]*models.Task, error) {
	const op = "repo.taskRepo.Tasks"

	query := `
		SELECT id, user_id, name, description, start_time, end_time, total_hours, created_at, updated_at
		FROM tasks
	`

	rows, err := r.db.Query(query)
	if err != nil {
		r.log.Debugf("%s %w", op, err)
		return nil, fmt.Errorf("%s %w", op, err)
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Name,
			&task.Description,
			&task.StartTime,
			&task.EndTime,
			&task.TotalHours,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			r.log.Debugf("%s %w", op, err)
			return nil, fmt.Errorf("%s %w", op, err)
		}
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		r.log.Debugf("%s %w", op, err)
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return tasks, nil
}

func (r *taskRepo) GetWorkloads(userID uint, startDate, endDate time.Time) ([]*models.Workload, error) {
	const op = "repo.taskRepo.GetWorkloads"

	query := `
		SELECT name, SUM(total_hours) AS total_hours
		FROM tasks
		WHERE user_id = $1 AND start_time >= $2 AND end_time <= $3
		GROUP BY name
		ORDER BY total_hours DESC
	`

	rows, err := r.db.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}
	defer rows.Close()

	workloads := make([]*models.Workload, 0)
	for rows.Next() {
		var workload models.Workload
		err := rows.Scan(&workload.TaskName, &workload.TotalHours)
		if err != nil {
			return nil, fmt.Errorf("%s %w", op, err)
		}
		workloads = append(workloads, &workload)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return workloads, nil
}
