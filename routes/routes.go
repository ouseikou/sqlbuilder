package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ouseikou/sqlbuilder/controller"
	"github.com/ouseikou/sqlbuilder/demo"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 不需要v1版本的公共服务的路由放外面
	// r.POST("/auth", api.GetAuth)

	apiV1 := router.Group("/api/v1")
	apiDemo := router.Group("/api/demo")

	{
		{
			// golang 练习
			apiDemo.GET("/database/:id", demo.GetDatabase)
			// curl -X GET 'localhost:7077/api/demo/sqlbuilder/gen?limit=10' | jq
			apiDemo.GET("/sqlbuilder/gen", demo.GetGenerateSql)
			// sudo apt install jq
			// curl -X POST localhost:7077/api/demo/sqlbuilder/product -d '{"category":""}' -H "Content-Type: application/json" -s | jq
			// curl -X POST localhost:7077/api/demo/sqlbuilder/product -d '{"category":"Doohickey"}' -H "Content-Type: application/json" -s | jq
			apiDemo.POST("/sqlbuilder/product", demo.GetGenProduct)
		}

		// 暴露接口
		apiV1.POST("/sqlbuilder", controller.GenerateHBQL)
		apiV1.POST("/sqlbuilder/proto", controller.GenerateHBQLByProto)
	}
}
