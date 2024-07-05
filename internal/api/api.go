package api

import (
	"test-task/common/http/middleware"
	"test-task/common/http/request"
	"test-task/infra"
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

}
