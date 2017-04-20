package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gnadlinger/Presentation/handler"
)

func CreateRoutes(){
	r:=gin.Default()
	v1:=r.Group("api/product")
	{
		v1.POST("/postproduct",handler.PostProduct)
		v1.GET("/getproducts",handler.GetAllProducts)
	}
	r.Run(":8080")
}
