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
		user := api.Group("/user")
		{
			user.GET("/", userHandler.GetAllUsers)
			user.GET("/filters", userHandler.UsersWithFiltersAndPagination)
			user.POST("/", userHandler.CreateUser)
			user.PATCH("/:id", userHandler.UpdateUser)
			user.DELETE("/:id", userHandler.DeleteUser)

			task := user.Group(":id/task")
			{
				task.POST("/:id/start", taskHandler.StartTask)
				task.POST("/:id/stop", taskHandler.StopTask)
			}
		}

		task := api.Group("/task")
		{
			task.GET("/", taskHandler.GetAllTasks)
			task.POST("/:id", taskHandler.CreateTask)
			task.PATCH("/:id", taskHandler.UpdateTask)
			task.DELETE("/:id", taskHandler.DeleteTask)
		}
	}

}
