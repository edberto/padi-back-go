package route

import (
	"padi-back-go/packages/condition"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	//initialize objects
	condition := condition.NewCondition()
	//register routes
	api := r.Group("")
	{
		api.GET("/condition/:label-id", condition.FindLabelHandler)
	}
}
