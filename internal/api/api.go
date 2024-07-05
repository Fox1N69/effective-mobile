package api

import (
	"test-task/common/http/middleware"
	"test-task/common/http/request"
	"test-task/infra"
	v1 "test-task/internal/api/v1"
	"test-task/internal/manager"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Server interface {
	Run()
}

type server struct {
	infra       infra.Infra
	gin         *gin.Engine
	service     manager.ServiceManager
	middleware  middleware.Middleware
	redisClient *redis.Client
}

func NewServer(infra infra.Infra, redisClient *redis.Client) Server {

	return &server{
		infra:       infra,
		gin:         gin.Default(),
		service:     manager.NewServiceManager(infra),
		middleware:  middleware.NewMiddleware(infra.Config().GetString("secret.key")),
		redisClient: redisClient,
	}
}

func (c *server) Run() {
	c.gin.Use(c.middleware.CORS())
	c.handlers()
	c.v1()

	c.gin.Run(c.infra.Port())
}

func (c *server) handlers() {
	h := request.DefaultHandler()

	c.gin.NoRoute(h.NoRoute)
	c.gin.GET("/", h.Index)
}

func (c *server) v1() {
	userHandler := v1.NewUserHandler(c.service.UserService())
	taskHandler := v1.NewTaskHandler(c.service.TaskService())

	api := c.gin.Group("/api")
	{
		api.Use(c.middleware.CORS())
		user := api.Group("/user")
		{
			user.GET("/", userHandler.GetAllUsers)
			user.GET("/filters", userHandler.UsersWithFiltersAndPagination)
			user.GET("/:id/workloads", taskHandler.GetWorkloads)
			user.POST("/", userHandler.CreateUser)
			user.PATCH("/:id", userHandler.UpdateUser)
			user.DELETE("/:id", userHandler.DeleteUser)

			// Вложенная группа для задач пользователя
			task := user.Group("/:user_id/task")
			{
				task.POST("/:task_id/start", taskHandler.StartTask)
				task.POST("/:task_id/stop", taskHandler.StopTask)
			}
		}

		// Роуты для задач
		tasks := api.Group("/task")
		{

			tasks.POST("/", taskHandler.CreateTask)
			tasks.GET("/", taskHandler.GetAllTasks)
			tasks.PATCH("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}
}
