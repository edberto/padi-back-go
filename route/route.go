package route

import (
	"padi-back-go/packages/condition"
	"padi-back-go/packages/register"
	"padi-back-go/packages/session"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	//initialize objects
	condition := condition.NewCondition()
	register := register.NewRegister()
	session := session.NewSession()

	//register routes
	api := r.Group("")
	{
		api.POST("/register", register.RegisterHandler)
		api.GET("/condition/:label-id", condition.FindLabelHandler)
		api.POST("/login", session.LoginHandler)

	}
}
