package route

import (
	"padi-back-go/packages/condition"
	"padi-back-go/packages/register"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	//initialize objects
	condition := condition.NewCondition()
	register := register.NewRegister()
	//register routes
	api := r.Group("")
	{
		api.POST("/register", register.RegisterHandler)
		api.GET("/condition/:label-id", condition.FindLabelHandler)

	}
}
