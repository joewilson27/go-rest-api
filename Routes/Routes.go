//Routes/Routes.go

package Routes

import (
	"go-rest-api/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grpl := r.Group("user-api")
	{
		grpl.GET("users", Controllers.GetUsers)
		grpl.POST("users", Controllers.CreateUser)
		grpl.GET("users/:id", Controllers.GetUserByID)
		grpl.PUT("users/:id", Controllers.UpdateUser)
		grpl.DELETE("users/:id", Controllers.DeleteUser)
	}
	return r
}
