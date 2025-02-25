package infrastructure

import (
	"github_wb/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(engines *gin.Engine) {

	routes := engines.Group("pull_request")

	{
		routes.POST("process", handlers.PullRequestEvent)
	}

}
