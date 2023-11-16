package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type College struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	CollegeRank uint   `json:"college_rank"`
	Code        uint   `json:"code"`
	Name        string `json:"name"`
	Grade       string `json:"grade"`
}

var db *gorm.DB

func main() {
	// 连接数据库
	dsn := "root:mm..1213@tcp(127.0.0.1:3306)/colleges"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err.Error())
	}

	// 自动迁移数据库表结构
	err = db.AutoMigrate(&College{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err.Error())
	}

	// 创建Gin路由
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	r.Use(cors.New(config))
	// 设置路由处理函数

	r.GET("/colleges", getCollegesHandler)
	r.POST("/colleges", addCollegeHandler)
	r.DELETE("/colleges/:rank", deleteCollegeHandler)

	// 启动HTTP服务器
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("HTTP服务器启动失败:", err.Error())
	}
}

func getCollegesHandler(c *gin.Context) {
	// 查询数据库获取学院数据
	var colleges []College
	result := db.Find(&colleges)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// 发送JSON响应
	c.JSON(http.StatusOK, colleges)
}

func addCollegeHandler(c *gin.Context) {
	// 解析请求中的JSON数据
	var college College
	if err := c.ShouldBindJSON(&college); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// 将学院数据添加到数据库
	result := db.Create(&college)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// 发送JSON响应
	c.JSON(http.StatusCreated, college)
}

func deleteCollegeHandler(c *gin.Context) {
	// 从URL参数中获取rank
	rank := c.Param("rank")

	// 从数据库中删除对应rank的学院数据
	result := db.Delete(&College{}, "college_rank = ?", rank)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// 发送成功响应
	c.JSON(http.StatusOK, gin.H{"message": "College deleted successfully"})
}
