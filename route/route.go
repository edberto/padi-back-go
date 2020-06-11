package route

import (
	"padi-back-go/middleware"
	"padi-back-go/packages/condition"
	"padi-back-go/packages/prediction"
	"padi-back-go/packages/register"
	"padi-back-go/packages/session"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	//initialize objects
	condition := condition.NewCondition()
	register := register.NewRegister()
	session := session.NewSession()
	prediction := prediction.NewPrediction()

	//middlewares
	authMW := middleware.SetTokenMiddleware()

	//register routes
	api := r.Group("")
	{
		api.POST("/register", register.RegisterHandler)
		api.POST("/login", session.LoginHandler)
		api.POST("/refresh", session.RefreshHandler)
		api.Use(authMW)
		api.GET("/condition/:label-id", condition.FindLabelHandler)
		api.DELETE("/logout", session.LogoutHandler)
		api.GET("/prediction", prediction.FindAllHandler)
		api.POST("/prediction", prediction.InsertOneHandler)

	}
}
