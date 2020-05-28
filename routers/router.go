/***************************************************
 ** @Desc : 注册相关路由
 ** @Time : 2020/5/28 13:56
 ** @Author : JiangFeng
 ** @File : error_gateway
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2020/5/28 16:56
 ** @Software: VScode
****************************************************/
package routers

import (
	"Gin_Blog/pkg/setting"
	v1 "Gin_Blog/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	// r.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })

	return r
}
