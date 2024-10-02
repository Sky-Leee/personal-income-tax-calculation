package main

import (
	"personal-income-tax-calculation-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("../static/index.html")
	
	r.GET("/tax", handlers.RenderTaxPage)
	r.POST("/api/calculate_tax", handlers.TaxCaculate)

	r.Run(":8080") // 启动服务
}
