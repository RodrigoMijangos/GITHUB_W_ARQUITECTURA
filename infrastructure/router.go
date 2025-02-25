package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {

	routes := engine.Group("pull_request")

	{
		routes.POST("process", PullRequestEvent)
	}

}
