package main

import (
	"github.com/forhsd/logger"
	"github.com/gin-gonic/gin"
	"github.com/ouseikou/sqlbuilder/config"
	"github.com/ouseikou/sqlbuilder/routes"
	"github.com/spf13/viper"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化数据库连接
	//db := repository.InitMysql()
	//db.AutoMigrate(&Model{}) // 根据 struct DDL建表

	//db := repository.InitPostgres()

	// 初始化 Redis 连接池
	//redisPool := repository.InitRedis()

	// 创建 Gin 实例
	router := gin.Default()

	// 注入数据库和 Redis 实例到 Gin 的中间件中
	router.Use(func(c *gin.Context) {
		//c.Set("db", db)
		//c.Set("redisPool", redisPool)
		c.Next()
	})

	// 配置路由
	routes.SetupRoutes(router)

	// 启动 HTTP 服务器
	port := viper.GetString("server.port")
	err := router.Run(":" + port)
	if err != nil {
		logger.Fatal("Failed to start server: %v", err)
	}
}
