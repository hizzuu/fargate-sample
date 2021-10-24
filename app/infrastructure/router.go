package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/hizzuu/app/interfaces/controllers"
	"github.com/hizzuu/app/interfaces/repository"
	"github.com/hizzuu/app/usecase/interactor"
)

type router struct {
	e *gin.Engine
}

func NewRouter() *router {
	if !AppConf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := &router{gin.New()}
	router.initMiddleware()
	router.initRouting()
	return router
}

func (r *router) Run() {
	r.e.Run(":" + ApiConf.Addr)
}

func (r *router) initMiddleware() {
	r.e.Use(gin.Logger())
	r.e.Use(gin.Recovery())
}

func (r *router) initRouting() {
	userCtrl := controllers.NewUserController(
		interactor.NewUserInteractor(
			repository.NewUserRepository(),
		),
	)

	v1 := r.e.Group("/v1")
	{
		v1.GET("/users/:id", func(ctx *gin.Context) { userCtrl.Get(ctx) })
		v1.GET("/users", func(ctx *gin.Context) { userCtrl.List(ctx) })
		v1.POST("/users", func(ctx *gin.Context) { userCtrl.Create(ctx) })
		v1.PUT("/users/:id", func(ctx *gin.Context) { userCtrl.Update(ctx) })
		v1.DELETE("/users/:id", func(ctx *gin.Context) { userCtrl.Delete(ctx) })
	}
}
