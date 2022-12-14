package routes

import (
	"github.com/Mersock/golang-clean-architecture-demo/databases"
	"github.com/Mersock/golang-clean-architecture-demo/deliveries"
	"github.com/Mersock/golang-clean-architecture-demo/repositories"
	"github.com/Mersock/golang-clean-architecture-demo/usecases"
	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter() *gin.Engine {

	todoRepo := repositories.NewToDoRepository(databases.DB)
	todoUseCase := usecases.NewToDoUseCase(todoRepo)
	todoHandler := deliveries.NewToDoHandler(todoUseCase)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", todoHandler.GetAllTodo)
		v1.POST("todo", todoHandler.CreateATodo)
		v1.GET("todo/:id", todoHandler.GetATodo)
		v1.PUT("todo/:id", todoHandler.UpdateATodo)
		v1.DELETE("todo/:id", todoHandler.DeleteATodo)
	}
	return r
}
