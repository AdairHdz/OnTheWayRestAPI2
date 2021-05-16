package states

import (	
	"github.com/gin-gonic/gin"	
)


func Routes(route *gin.RouterGroup) {
	states := route.Group("/states")
	{
		states.GET("/", func(context *gin.Context){
			
		})

		states.GET("/:stateId/cities", func(context *gin.Context){
			
		})
	}
}